package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	apiURL = "https://api.github.com/users/%s/events"
)

func (e *GitHubErrorResponse) Error() string {
	return fmt.Sprintf("GitHub API error: %s (Status: %s)", e.Message, e.Status)
}

func fetchGitHubActivity(ctx context.Context, username string) ([]Event, error) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	url := fmt.Sprintf(apiURL, username)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			fmt.Println("Error closing response body:", err)
		}
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		var githubErr GitHubErrorResponse
		if err := json.Unmarshal(body, &githubErr); err != nil {
			return nil, fmt.Errorf("unexpected response status: %d", resp.StatusCode)
		}
		switch githubErr.Status {
		case "404":
			return nil, fmt.Errorf("User '%s' not found on GitHub", username)
		case "403":
			return nil, fmt.Errorf("Access to GitHub API is forbidden. You may have exceeded the rate limit")
		default:
			return nil, &githubErr
		}
	}

	var events []Event
	err = json.Unmarshal(body, &events)
	if err != nil {
		return nil, err
	}

	return events, nil
}

// UnmarshalJSON is a custom unmarshaler for the Event type.
func (e *Event) UnmarshalJSON(data []byte) error {
	type Alias Event
	aux := &struct {
		*Alias
		CreatedAt string `json:"created_at"`
	}{
		Alias: (*Alias)(e),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	parsedTime, err := time.Parse(time.RFC3339, aux.CreatedAt)
	if err != nil {
		return err
	}
	e.CreatedAt = parsedTime
	return nil
}
