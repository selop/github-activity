package main

import (
	"fmt"
	"strings"
	"time"
)

// timeAgo returns a string with the time difference between the current time and the given time
func timeAgo(t time.Time) string {
	duration := time.Since(t)
	switch {
	case duration.Hours() > 24*30:
		return fmt.Sprintf("%d months ago -", int(duration.Hours()/24/30))
	case duration.Hours() > 24:
		return fmt.Sprintf("%d days ago -", int(duration.Hours()/24))
	case duration.Hours() > 1:
		return fmt.Sprintf("%d hours ago -", int(duration.Hours()))
	case duration.Minutes() > 1:
		return fmt.Sprintf("%d minutes ago -", int(duration.Minutes()))
	default:
		return "just now"
	}
}

func capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

// shortenURL shortens the URL to the last two parts of the URL
func shortenURL(url string) string {
	parts := strings.Split(url, "/")
	if len(parts) >= 2 {
		return parts[len(parts)-2] + "/" + parts[len(parts)-1][:7]
	}
	return url
}
