package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/dmandevv/github-activity/internal/api"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		printHelp()
		return
	}
	username := args[1]
	perPage := flag.Int("per_page", 15, "Number of events to fetch per page (max 100)")
	page := flag.Int("page", 1, "Page number to fetch")
	help := flag.Bool("help", false, "Show help")
	flag.CommandLine.Parse(os.Args[2:])

	if *help {
		printHelp()
		return
	}

	url := fmt.Sprintf("https://api.github.com/users/%s/events?per_page=%d&page=%d", username, *perPage, *page)
	resp, err := http.Get(url)

	if resp.StatusCode == 404 {
		fmt.Println("User not found:", username)
		return
	}
	if resp.StatusCode == 503 {
		fmt.Println("GitHub API is currently unavailable. Please try again later.")
		return
	}

	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}
	printActivity(body)
}

func printHelp() {
	fmt.Println("Usage: github-activity <GITHUB_USERNAME> -per_page <NUMBER> -page <NUMBER>")
}

func printActivity(data []byte) {
	var events []api.GitHubEvent
	err := json.Unmarshal(data, &events)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	for _, event := range events {
		fmt.Println(event.Type, "at", event.CreatedAt.Local(), "in repo", event.Repo.Name)
	}
}
