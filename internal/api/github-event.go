package api

import "time"

type GitHubEvent struct {
	ID        string      `json:"id"`
	Type      string      `json:"type"`
	Actor     Actor       `json:"actor"`
	Repo      Repo        `json:"repo"`
	Payload   interface{} `json:"payload"`
	Public    bool        `json:"public"`
	CreatedAt time.Time   `json:"created_at"`
}

type Actor struct {
	ID        int    `json:"id"`
	Login     string `json:"login"`
	AvatarURL string `json:"avatar_url"`
	URL       string `json:"url"`
}

type Repo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}
