package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/google/go-github/v43/github"
	"golang.org/x/oauth2"
)

var pr_actions = map[string]string{
	"approve": "APPROVE",
	"request": "CHANGES_REQUEST",
	"comment": "COMMENT",
}

func print(color string, message string) {
	var color_palettes = map[string]string{
		"red":    "\033[31m",
		"green":  "\033[32m",
		"yellow": "\033[33m",
		"reset":  "\033[0m",
	}

	fmt.Println(string(color_palettes[color]), message, color_palettes["reset"])
}

func main() {
	var pr_url = flag.String("url", "", "PR URL")
	var message = flag.String("message", "LGTM", "message body")
	var action = flag.String("action", "approve", "review action")
	var help = flag.Bool("help", false, "Help message")
	flag.Parse()

	if *help {
		fmt.Println("Usage: goapprove [options]")
		print("yellow", "Options:")
		fmt.Println("  -url: PR URL")
		fmt.Println("  -message: message body")
		fmt.Println("  -action: review action")
		fmt.Println("  -help: help message")
		fmt.Println("")
		print("yellow", "Requirements:")
		fmt.Println("GITHUB_TOKEN environment variable is needed")
		fmt.Println("Generate token at https://github.com/settings/tokens/new?scopes=repo&description=goapprove-cli")
		fmt.Println("")
		print("yellow", "Example:")
		fmt.Println("goapprove -url https://github.com/amazingandyyy/goapprove/pull/1 -action comment -message \"LGTM 🚀\"")
		os.Exit(0)
	}

	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		print("yellow", "Generate token at https://github.com/settings/tokens/new?scopes=repo&description=goapprove-cli")
		fmt.Print("GITHUB_TOKEN: ")
		fmt.Scanf("%s", &token)
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
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	body := fmt.Sprintf("%s\n\n%v", *message, "<p align='right'>☘️ reviewed from my terminal via <a href='https://github.com/amazingandyyy/goapprove' target='_blank'>goapprove</a></p>")

	opts := &github.PullRequestReviewRequest{
		Body:  github.String(body),
		Event: github.String(pr_actions[*action]),
	}
	pr_number, _ := strconv.Atoi(pr[4])
	pr_review, resp, _ := client.PullRequests.CreateReview(ctx, pr[1], pr[2], pr_number, opts)

	defer resp.Body.Close()

	type GitHubResponse struct {
		Errors []string `json:"errors,omitempty"`
	}
	var j GitHubResponse
	json.NewDecoder(resp.Body).Decode(&j)
	if len(j.Errors) == 0 && resp.StatusCode != 200 {
		print("red", fmt.Sprintf("Error: %v", j.Errors[0]))
		os.Exit(1)
	}

	done_msg := fmt.Sprintf("%ved with message on %q %v", *action, *message, *pr_review.HTMLURL)
	print("green", done_msg)
}
