# GitHub Activity Viewer

## Description

GitHub Activity Viewer is a command-line tool that fetches and displays recent GitHub activity for a specified user. It provides a quick and easy way to see a user's latest actions on GitHub, including push events, issue interactions, and repository stars.

## Features

- Fetch recent GitHub activity for any public user
- Display activity in a human-readable format
- Show timestamps in a "time ago" format
- Highlight different types of events with emojis

## Usage

Run the program with a GitHub username as an argument:
```
./github-activity <username>
```

### Sample output

```
./github-activity selop  
1 days ago - 🔨 Pushed commits to selop/tasktracker [commits/fc45c17]
1 days ago - 🔨 Pushed commits to selop/tasktracker [commits/425bed3]
1 days ago - 🔨 Pushed commits to selop/tasktracker [commits/7dd8c2c]
1 days ago - 🔨 Pushed commits to selop/tasktracker [commits/715cfeb]
1 days ago - 🔨 Pushed commits to selop/tasktracker [commits/a086012]
1 days ago - 🔨 Pushed commits to selop/tasktracker [commits/418e705]
1 days ago - 🔔 CreateEvent on selop/tasktracker
1 days ago - 🔔 CreateEvent on selop/tasktracker
1 months ago - ⭐ Starred streetpea/chiaki-ng
1 months ago - 🔔 IssueCommentEvent on inkle/ink
1 months ago - ⭐ Starred Ocean-Moist/FizzBuzz
1 months ago - 🔔 PullRequestEvent on PacktPublishing/Learning-Spring-Boot-3.0-Third-Edition
1 months ago - 🔔 CreateEvent on selop/Learning-Spring-Boot-3.0-Third-Edition
1 months ago - 🔔 ForkEvent on PacktPublishing/Learning-Spring-Boot-3.0-Third-Edition
1 months ago - 🔔 IssueCommentEvent on PacktPublishing/Learning-Spring-Boot-3.0-Third-Edition
2 months ago - 🐛 Closed an issue in wimdeblauwe/modern-frontends-with-htmx-sources
2 months ago - 🔔 IssueCommentEvent on wimdeblauwe/modern-frontends-with-htmx-sources
2 months ago - 🐛 Closed an issue in wimdeblauwe/modern-frontends-with-htmx-sources
```

## Installation

To install GitHub Activity Viewer, make sure you have Go installed on your system (version 1.22.3 or later), then follow these steps:

1. Clone the repository:
   ```
   git clone https://github.com/selop/github-activity.git
   ```

2. Navigate to the project directory:
   ```
   cd github-activity
   ```

3. Build the project:
   ```
   go build
   ```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Thanks to GitHub for providing access to user activity data.
- This project was partly composed with Cursor AI and claude-3.5-sonnet.