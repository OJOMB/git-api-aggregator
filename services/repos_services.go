package services

import (
	"sync"

	"github.com/OJOMB/git-api-aggregator/config"
	"github.com/OJOMB/git-api-aggregator/domain/github"
	"github.com/OJOMB/git-api-aggregator/domain/gitrepositories"
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
	CreateRepos(req *gitrepositories.CreateReposRequest) *gitrepositories.CreateReposResponse
}

type repos struct{}

func (rs *repos) CreateRepo(req *gitrepositories.CreateRepoRequest) (*gitrepositories.CreateRepoResponse, utils.APIError) {
	validationErr := req.Validate()
	if validationErr != nil {
		return nil, validationErr
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

	return &gitrepositories.CreateRepoResponse{
		ID:    resp.ID,
		Name:  resp.Name,
		Owner: resp.Owner.Login,
	}, nil
}

func (rs *repos) CreateRepos(req *gitrepositories.CreateReposRequest) *gitrepositories.CreateReposResponse {
	resultsChan := make(chan *gitrepositories.CreateRepoResult)

	var wg sync.WaitGroup
	for _, r := range req.Requests {
		go func(request gitrepositories.CreateRepoRequest, resultsChan chan *gitrepositories.CreateRepoResult) {
			result := gitrepositories.NewCreateRepoResult(rs.CreateRepo(&request))
			resultsChan <- &result
			wg.Done()
		}(r, resultsChan)
		wg.Add(1)
	}

	wg.Wait()
	close(resultsChan)

	var createReposResponse gitrepositories.CreateReposResponse
	for result := range resultsChan {
		createReposResponse.Results = append(createReposResponse.Results, *result)
	}

	return &createReposResponse
}
