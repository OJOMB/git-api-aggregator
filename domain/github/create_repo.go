package github

import "github.com/gin-gonic/gin"

// Repos is the exported entrypoint to the Github Repos world
var Repos ReposInterface

func init() {
	Repos = &repos{}
}

// ReposInterface is the interface for the Github Repos world
type ReposInterface interface{}

// CreateRepoRequest models the structure of the request for a Create Repo API call
type CreateRepoRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Homepage    string `json:"homepage"`
	Private     bool   `json:"private"`
	HasIssues   bool   `json:"has_issues"`
	HasProjects bool   `json:"has_projects"`
	HasWiki     bool   `json:"has_wiki"`
}

// CreateRepoResponse models the structure of the response for a Create Repo API call
type CreateRepoResponse struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	FullName   string `json:"full_name"`
	Owner      Owner  `json:"owner"`
	Permission string `json:"permission"`
}

// Owner represents a repository owner.
type Owner struct {
	ID      int    `json:"id"`
	Login   string `json:"login"`
	URL     string `json:"url"`
	HTMLURL string `json:"html_url"`
}

// RepoPermissions models the structure of a Repo Permissions
type RepoPermissions struct {
	IsAdmin bool `json:"is_admin"`
	HasPull bool `json:"has_pull"`
	HasPush bool `json:"has_push"`
	Admin   bool `json:"admin"`
	Push    bool `json:"push"`
	Pull    bool `json:"pull"`
}

type repos struct{}

func (g *repos) HandleCreateRepo() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
