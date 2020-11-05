package services

import (
	"github.com/OJOMB/git-api-aggregator/domain/github"
	"github.com/OJOMB/git-api-aggregator/domain/gitusers"
	"github.com/OJOMB/git-api-aggregator/providers/githubprovider"
	"github.com/OJOMB/git-api-aggregator/utils"
)

// Users is the swappable interface instance
var Users UsersServiceInterface

func init() {
	Users = &users{}
}

// UsersServiceInterface is the interface for the Repo Service
type UsersServiceInterface interface {
	GetUser(getUserRequest *gitusers.GetUserRequest) (*gitusers.GetUserResponse, utils.APIError)
}

type users struct{}

func (u *users) GetUser(getUserRequest *gitusers.GetUserRequest) (*gitusers.GetUserResponse, utils.APIError) {
	resp, errResp := githubprovider.GetUser(&github.GetUserRequest{Login: getUserRequest.UserName})
	if errResp != nil {
		return nil, utils.NewAPIError(errResp.StatusCode, errResp.Message)
	}

	return &gitusers.GetUserResponse{
		ID:          resp.ID,
		UserName:    resp.Login,
		URL:         resp.URL,
		HTMLURL:     resp.HTMLURL,
		Location:    resp.Location,
		Name:        resp.Name,
		Blog:        resp.Blog,
		Bio:         resp.Bio,
		Followers:   resp.Followers,
		Following:   resp.Following,
		PublicRepos: resp.PublicRepos,
		CreatedAt:   resp.CreatedAt,
	}, nil
}
