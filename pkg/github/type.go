package github

import "time"

type PushEvent struct {
	Ref        string     `json:"ref,omitempty"`
	Before     string     `json:"before,omitempty"`
	After      string     `json:"after,omitempty"`
	Created    bool       `json:"created,omitempty"`
	Deleted    bool       `json:"deleted,omitempty"`
	Forced     bool       `json:"forced,omitempty"`
	BaseRef    *string    `json:"base_ref,omitempty"`
	Compare    string     `json:"compare,omitempty"`
	Repository Repository `json:"repository,omitempty"`
	Pusher     Pusher     `json:"pusher,omitempty"`
	Sender     Sender     `json:"sender,omitempty"`
	Commits    []Commit   `json:"commits,omitempty"`
	HeadCommit Commit     `json:"head_commit,omitempty"`
}

type Repository struct {
	ID            int64     `json:"id,omitempty"`
	NodeID        string    `json:"node_id,omitempty"`
	Name          string    `json:"name,omitempty"`
	FullName      string    `json:"full_name,omitempty"`
	Private       bool      `json:"private,omitempty"`
	Owner         Owner     `json:"owner,omitempty"`
	HTMLURL       string    `json:"html_url,omitempty"`
	Description   *string   `json:"description,omitempty"`
	Fork          bool      `json:"fork,omitempty"`
	URL           string    `json:"url,omitempty"`
	CreatedAt     int64     `json:"created_at,omitempty"`
	UpdatedAt     time.Time `json:"updated_at,omitempty"`
	PushedAt      int64     `json:"pushed_at,omitempty"`
	GitURL        string    `json:"git_url,omitempty"`
	SSHURL        string    `json:"sshurl,omitempty"`
	CloneURL      string    `json:"clone_url,omitempty"`
	Size          int       `json:"size,omitempty"`
	DefaultBranch string    `json:"default_branch,omitempty"`
	MasterBranch  string    `json:"master_branch,omitempty"`
}

type Owner struct {
	Name   string `json:"name,omitempty"`
	Email  string `json:"email,omitempty"`
	Login  string `json:"login,omitempty"`
	ID     int64  `json:"id,omitempty"`
	NodeID string `json:"node_id,omitempty"`
	Type   string `json:"type,omitempty"`
}

type Pusher struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

type Sender struct {
	Login  string `json:"login,omitempty"`
	ID     int64  `json:"id,omitempty"`
	NodeID string `json:"node_id,omitempty"`
	Type   string `json:"type,omitempty"`
}

type Commit struct {
	ID        string    `json:"id,omitempty"`
	TreeID    string    `json:"tree_id,omitempty"`
	Distinct  bool      `json:"distinct,omitempty"`
	Message   string    `json:"message,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
	URL       string    `json:"url,omitempty"`
	Author    Author    `json:"author,omitempty"`
	Committer Author    `json:"committer,omitempty"`
	Added     []string  `json:"added,omitempty"`
	Removed   []string  `json:"removed,omitempty"`
	Modified  []string  `json:"modified,omitempty"`
}

type Author struct {
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Username string `json:"username,omitempty"`
}
