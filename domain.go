package main

import (
	"time"
)

// Event represents a GitHub event.
type Event struct {
	Type      string    `json:"type"`
	Repo      Repo      `json:"repo"`
	CreatedAt time.Time `json:"created_at"`
	Payload   struct {
		Action  string `json:"action"`
		Commits []struct {
			URL string `json:"url"`
		} `json:"commits"`
	} `json:"payload"`
}

// Repo represents a GitHub repository.
type Repo struct {
	Name string `json:"name"`
}

// GitHubErrorResponse represents an error response from the GitHub API.
type GitHubErrorResponse struct {
	Message          string `json:"message"`
	DocumentationURL string `json:"documentation_url"`
	Status           string `json:"status"`
}
