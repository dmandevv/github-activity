# GitHub Activity CLI 🚀

A small command-line tool written in Go that fetches and summarizes a GitHub user's public activity events (pushes, PRs, issues, etc.). This project is an exercise from the roadmap.sh project: https://roadmap.sh/projects/github-user-activity 📚

**What it does**

- **Fetches** public events for a GitHub user
- **Summarizes** common event types (pushes, pull requests, issues, stars)
- **Outputs** results to stdout so you can pipe or format them as you like

**Prerequisites**

- Go 1.25.1 or newer installed
- Network access to GitHub's public API

**Install**

There are a few easy ways to get and run this project.

- **Install using Go**

```shell
go install github.com/dmandevv/github-activity@latest
github-activity dmandevv
```
This installs the binary into your `$GOBIN` (or `$GOPATH/bin`) so you can run it globally.

- **Clone and build from source**

```shell
git clone https://github.com/dmandevv/github-activity.git
cd github-activity
go build -o github-activity .
```

After building you can run the generated `github-activity` binary.

- **Run directly (quick, no install)**

```shell
go run main.go dmandevv
```

## Usage examples ✅

Below are common ways to run the CLI and combine flags. Replace `dmandevv` with the GitHub username you want to inspect.

- Show usage:

```shell
github-activity dmandevv -help
```

- Show more events per page (up to 100):

```shell
# fetch up to 50 events in the first page
github-activity dmandevv -per_page 50
```

- Paginate through results (page 2):

```shell
# fetch the second page of results with 30 events per page
github-activity dmandevv -per_page 30 -page 2
```

- User not found:

```shell
github-activity some-nonexistent-user-12345
# => "User not found: some-nonexistent-user-12345"
```

**Notes**

- This repository is an educational implementation aligned with the roadmap.sh project: https://roadmap.sh/projects/github-user-activity
- If you plan to make many requests or need more data, consider using a GitHub API token to increase rate limits (not implemented by default).
