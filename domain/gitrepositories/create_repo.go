package gitrepositories

import (
	"strings"

	"github.com/OJOMB/git-api-aggregator/utils"
)

// CreateRepoRequest d efines the expected data structure of a create repo request
type CreateRepoRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Private     bool   `json:"private"`
}

// Validate checks that the values we receive as CreateRepoRequest fields are valid
func (req *CreateRepoRequest) Validate() utils.APIError {
	req.Name = strings.TrimSpace(req.Name)
	if req.Name == "" {
		return utils.NewBadRequestError("Repo `name` is a required field")
	}
	return nil
}

// CreateRepoResponse defines the expected data structure of a create repo response
type CreateRepoResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Owner string `json:"owner"`
}

// CreateReposRequest defines the expected data structure of a create repos response
type CreateReposRequest struct {
	Requests []CreateRepoRequest `json:"requests"`
}

// CreateReposResponse defines the expected data structure of a create repos response
type CreateReposResponse struct {
	Results []CreateRepoResult `json:"results"`
}

// CreateRepoResult stores the result of an individual CreateRepo request
type CreateRepoResult struct {
	Response *CreateRepoResponse `json:"response"`
	Error    utils.APIError      `json:"error"`
}

// NewCreateRepoResult returns a new CreateRepoResult
func NewCreateRepoResult(resp *CreateRepoResponse, err utils.APIError) CreateRepoResult {
	return CreateRepoResult{Response: resp, Error: err}
}
