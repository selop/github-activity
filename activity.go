// Package main provides functionality to fetch and display GitHub user activity.
package main

import (
	"fmt"
)

const (
	pushEvent  = "PushEvent"
	issueEvent = "IssuesEvent"
	watchEvent = "WatchEvent"
)

func displayActivity(events []Event) {
	for _, event := range events {
		timeAgo := timeAgo(event.CreatedAt)
		var link string
		switch event.Type {
		case pushEvent:
			if len(event.Payload.Commits) > 0 {
				shortURL := shortenURL(event.Payload.Commits[0].URL)
				link = fmt.Sprintf("[%s]", shortURL)
			}
			fmt.Printf("%s 🔨 Pushed commits to %s %s\n", timeAgo, event.Repo.Name, link)
		case issueEvent:
			fmt.Printf("%s 🐛 %s an issue in %s\n", timeAgo, capitalize(event.Payload.Action), event.Repo.Name)
		case watchEvent:
			fmt.Printf("%s ⭐ Starred %s\n", timeAgo, event.Repo.Name)
		default:
			fmt.Printf("%s 🔔 %s on %s\n", timeAgo, event.Type, event.Repo.Name)
		}
	}
}
