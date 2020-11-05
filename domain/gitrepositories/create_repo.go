package gitrepositories

// CreateRepoRequest models a create repo request
type CreateRepoRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Private     bool   `json:"private"`
}

// CreateRepoResponse models a create repo response
type CreateRepoResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Owner string `json:"owner"`
}
