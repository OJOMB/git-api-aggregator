package services

import (
	"strings"

	"github.com/OJOMB/git-api-aggregator/config"
	"github.com/OJOMB/git-api-aggregator/domain/github"
	"github.com/OJOMB/git-api-aggregator/domain/gitrepositories"
	"github.com/OJOMB/git-api-aggregator/domain/repositories"
	"github.com/OJOMB/git-api-aggregator/providers/githubprovider"
	"github.com/OJOMB/git-api-aggregator/utils"
)

// Repos is the swappable interface instance
var Repos ReposServiceInterface

func init() {
	Repos = &repos{}
}

// ReposServiceInterface is the interface for the Repo Service
type ReposServiceInterface interface {
	CreateRepo(req *gitrepositories.CreateRepoRequest) (*gitrepositories.CreateRepoResponse, utils.APIError)
}

type repos struct{}

func (rs *repos) CreateRepo(req *gitrepositories.CreateRepoRequest) (*gitrepositories.CreateRepoResponse, utils.APIError) {
	req.Name = strings.TrimSpace(req.Name)
	if req.Name == "" {
		return nil, utils.NewBadRequestError("Repo `name` is a required field")
	}

	createRepoRequest := &github.CreateRepoRequest{
		Name:        req.Name,
		Description: req.Description,
		Private:     req.Private,
	}

	resp, err := githubprovider.CreateRepo(createRepoRequest, "OJOMB", config.GetGithubAccessToken())
	if err != nil {
		return nil, utils.NewAPIError(err.StatusCode, err.Message)
	}

	return &repositories.CreateRepoResponse{
		ID:    resp.ID,
		Name:  resp.Name,
		Owner: resp.Owner.Login,
	}, nil
}
