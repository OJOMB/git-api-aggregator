package github

// GetUserRequest is the data model for a Github Get User Response
type GetUserRequest struct {
	Login string `json:"login"`
}

// GetUserResponse is the data model for a Github Get User Response
type GetUserResponse struct {
	ID          int    `json:"id"`
	Login       string `json:"login"`
	URL         string `json:"url"`
	HTMLURL     string `json:"html_url"`
	Location    string `json:"location"`
	Name        string `json:"name"`
	Blog        string `json:"blog"`
	Bio         string `json:"bio"`
	Followers   int    `json:"followers"`
	Following   int    `json:"following"`
	PublicRepos int    `json:"public_repos"`
	CreatedAt   string `json:"created_at"`
}
