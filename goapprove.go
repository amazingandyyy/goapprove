package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/google/go-github/v43/github"
	"golang.org/x/oauth2"
)

var (
	pr_url = flag.String("url", "", "PR URL")
)

func main() {

	flag.Parse()
	token := os.Getenv("GITHUB_AUTH_TOKEN")
	if token == "" {
		fmt.Println("Visit https://github.com/settings/tokens/new to generate a personal access token with [repo] scope")
		fmt.Print("GITHUB_AUTH_TOKEN: ")
		fmt.Scanf("%s", &token)
	}
	if *pr_url == "" {
		fmt.Print("PR URL: ")
		fmt.Scanf("%s", pr_url)
	}

	pr := strings.Split(*pr_url, "/")

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	opts := &github.PullRequestReviewRequest{
		Body:  github.String("LGTM"),
		Event: github.String("APPROVE"),
	}
	pr_number, _ := strconv.Atoi(pr[6])
	pr_review, _, err := client.PullRequests.CreateReview(ctx, pr[3], pr[4], pr_number, opts)

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Printf("%v is approved\n", *pr_review.HTMLURL)
}
