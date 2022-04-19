package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/google/go-github/v43/github"
	"golang.org/x/oauth2"
)

var goapproveConfigFilePath string = os.ExpandEnv("$HOME/.goapprove.json")

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var pr_actions = map[string]string{
	"approve": "APPROVE",
	"request": "CHANGES_REQUEST",
	"comment": "COMMENT",
}

func print(color string, message string) {
	color = strings.ToLower(color)
	var color_palettes = map[string]string{
		"red":    "\033[31m",
		"green":  "\033[32m",
		"yellow": "\033[33m",
		"reset":  "\033[0m",
	}

	fmt.Println(string(color_palettes[color]), message, color_palettes["reset"])
}

type GoapproveConfig struct {
	Ghtoken string `json:"github_token"`
}

func validateGhToken(ghToken string) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: ghToken})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	_, _, err := client.Repositories.List(ctx, "github", nil)
	if err != nil {
		print("red", "Invalid token")
		os.Exit(1)
	}
}

func renewGhToken() {
	ghToken := ""
	hasConfigFile := false
	print("yellow", "Generate token at https://github.com/settings/tokens/new?scopes=repo&description=goapprove-cli")
	fmt.Print("GITHUB_TOKEN: ")
	fmt.Scanf("%s", &ghToken)
	validateGhToken(ghToken)
	print("green", "Token is valid")

	var jsonBlob = []byte(`{"github_token": "` + ghToken + `"}`)
	goapproveConfig := GoapproveConfig{}
	err := json.Unmarshal(jsonBlob, &goapproveConfig)
	if err != nil {
		panic(err)
	}
	if _, err := os.Stat(goapproveConfigFilePath); err == nil {
		hasConfigFile = true
	}
	if os.Getenv("CI") != "true" {
		goapproveConfigJson, _ := json.Marshal(goapproveConfig)
		if !hasConfigFile {
			// create one
			f, err := os.Create(goapproveConfigFilePath)
			check(err)
			defer f.Close()
			err = ioutil.WriteFile(goapproveConfigFilePath, goapproveConfigJson, 0644)
			check(err)
		}
	}
}

func main() {
	var pr_url = flag.String("url", "", "PR URL")
	var message = flag.String("message", "LGTM", "message body")
	var action = flag.String("action", "approve", "review action")
	var auth = flag.Bool("auth", false, "renew github token")
	flag.Parse()

	if *auth {
		renewGhToken()
		os.Exit(0)
	}
	ghToken := os.Getenv("GITHUB_TOKEN")
	if _, err := os.Stat(goapproveConfigFilePath); err == nil {
		// config file does exist
		config, _ := os.ReadFile(goapproveConfigFilePath)
		var goapproveConfig GoapproveConfig
		err := json.Unmarshal(config, &goapproveConfig)
		if err != nil {
			panic(err)
		}
		if goapproveConfig.Ghtoken != "" {
			ghToken = goapproveConfig.Ghtoken
		}
		validateGhToken(ghToken)
	}

	if ghToken == "" {
		renewGhToken()
	}

	if *pr_url == "" {
		fmt.Print("PR URL: ")
		fmt.Scanf("%s", pr_url)
	}

	parsed_url, err := url.Parse(*pr_url)
	if err != nil {
		panic(err)
	}
	pr := strings.Split(parsed_url.Path, "/")

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: ghToken})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	body := fmt.Sprintf("%s\n\n%v", *message, "<p align='right'>☘️ reviewed via <a href='https://github.com/amazingandyyy/goapprove' target='_blank'>goapprove</a></p>")

	opts := &github.PullRequestReviewRequest{
		Body:  github.String(body),
		Event: github.String(pr_actions[*action]),
	}
	pr_number, _ := strconv.Atoi(pr[4])
	pr_owner := pr[1]
	pr_repo := pr[2]
	pr_review, resp, _ := client.PullRequests.CreateReview(ctx, pr_owner, pr_repo, pr_number, opts)

	defer resp.Body.Close()

	type GitHubResponse struct {
		Errors []string `json:"errors,omitempty"`
	}
	var j GitHubResponse
	err = json.NewDecoder(resp.Body).Decode(&j)
	check(err)
	if len(j.Errors) == 0 && resp.StatusCode != 200 {
		print("red", fmt.Sprintf("Error: %v", j.Errors))
		os.Exit(1)
	}

	done_msg := fmt.Sprintf("%v with message on %q %v", *action, *message, *pr_review.HTMLURL)
	print("green", done_msg)
}
