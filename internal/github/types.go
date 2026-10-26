package github

import "time"

// User is the subset of a GitHub user used by Okto.
type User struct {
	Login   string `json:"login"`
	ID      int64  `json:"id"`
	HTMLURL string `json:"html_url"`
}

// License is the subset of a repository license used by Okto.
type License struct {
	SPDXID string `json:"spdx_id"`
}

// Repository is the subset of the repository response used by Okto.
type Repository struct {
	ID               int64      `json:"id"`
	Name             string     `json:"name"`
	FullName         string     `json:"full_name"`
	Owner            User       `json:"owner"`
	HTMLURL          string     `json:"html_url"`
	Description      *string    `json:"description"`
	DefaultBranch    string     `json:"default_branch"`
	Language         *string    `json:"language"`
	License          *License   `json:"license"`
	StargazersCount  int        `json:"stargazers_count"`
	ForksCount       int        `json:"forks_count"`
	SubscribersCount int        `json:"subscribers_count"`
	OpenIssuesCount  int        `json:"open_issues_count"`
	Archived         bool       `json:"archived"`
	Private          bool       `json:"private"`
	CreatedAt        time.Time  `json:"created_at"`
	PushedAt         *time.Time `json:"pushed_at"`
}

// WeeklyCommit is one week of commit activity.
type WeeklyCommit struct {
	Days  [7]int `json:"days"`
	Total int    `json:"total"`
	Week  int64  `json:"week"`
}

// Release is the subset of a GitHub release used by Okto.
type Release struct {
	ID          int64      `json:"id"`
	TagName     string     `json:"tag_name"`
	Name        string     `json:"name"`
	Draft       bool       `json:"draft"`
	Prerelease  bool       `json:"prerelease"`
	CreatedAt   *time.Time `json:"created_at"`
	PublishedAt *time.Time `json:"published_at"`
	HTMLURL     string     `json:"html_url"`
}

// Contributor is the subset of a contributors response used by Okto.
type Contributor struct {
	Login         string `json:"login"`
	ID            int64  `json:"id"`
	HTMLURL       string `json:"html_url"`
	Contributions int    `json:"contributions"`
}

// GitCommit is the nested commit metadata.
type GitCommit struct {
	Author struct {
		Name  string    `json:"name"`
		Email string    `json:"email"`
		Date  time.Time `json:"date"`
	} `json:"author"`
	Message string `json:"message"`
}

// Commit is the subset of a commits response used by Okto.
type Commit struct {
	SHA     string    `json:"sha"`
	HTMLURL string    `json:"html_url"`
	Commit  GitCommit `json:"commit"`
	Author  *User     `json:"author"`
}

// RateLimit is the latest rate-limit state observed from GitHub headers.
type RateLimit struct {
	Limit     int
	Remaining int
	ResetAt   time.Time
	Resource  string
}
