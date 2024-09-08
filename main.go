package main

import (
	"context"
	"fmt"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	if len(os.Args) < 2 {
		return fmt.Errorf("Please provide a GitHub username")
	}

	username := os.Args[1]
	events, err := fetchGitHubActivity(context.Background(), username)
	if err != nil {
		return err
	}

	if len(events) == 0 {
		return fmt.Errorf("No recent activity found for user '%s'", username)
	}

	displayActivity(events)
	return nil
}
