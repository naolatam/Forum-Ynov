package dtos

type GitHubUserInfo struct {
	ID        int64  `json:"id"`
	NodeID    string `json:"node_id"`
	Login     string `json:"login"`
	Name      string `json:"name"`
	Emails    []GithubUserEmail
	AvatarURL string `json:"avatar_url"`
	Bio       string `json:"bio"`
}

type GithubUserEmail struct {
	Email    string       `json:"email"`
	Primary  bool         `json:"primary"`
	Verified *interface{} `json:"visibility"`
}
