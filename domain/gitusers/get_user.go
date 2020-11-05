package gitusers

// GetUserRequest models a create repo request
type GetUserRequest struct {
	UserName string `json:"user_name"`
}

// GetUserResponse models a create repo response
type GetUserResponse struct {
	ID          int    `json:"id"`
	UserName    string `json:"username"`
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
