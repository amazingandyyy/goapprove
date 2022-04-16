package main

import (
	"flag"
	"fmt"
)

type GitUser struct {
	Username string
}

func main() {
	pr_url := flag.String("l", "", "PR URL")
	flag.Parse()
	fmt.Println("PR URL:", *pr_url)
}
