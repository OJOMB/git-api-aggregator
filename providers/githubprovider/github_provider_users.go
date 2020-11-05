package githubprovider

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/OJOMB/git-api-aggregator/clients/rest"
	"github.com/OJOMB/git-api-aggregator/domain/github"
)

const apiGetUserFormat = "http://api.github.com/users/%s"

// GetUser performs a GET request to the users API for the given username
func GetUser(req *github.GetUserRequest) (*github.GetUserResponse, *github.ErrorResponse) {
	// construct request
	headers := &http.Header{}
	headers.Set("Authorization", GetAuthorizationHeader(token))
	endpoint := fmt.Sprintf(apiGetUserFormat, req.Login)
	resp, err := rest.Client.GET(endpoint, headers)

	respBodyJSON, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, &github.ErrorResponse{
			Message:    err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	if resp.StatusCode != 200 {
		var githubErrResp github.ErrorResponse
		err := json.Unmarshal(respBodyJSON, &githubErrResp)
		if err != nil {
			return nil, &github.ErrorResponse{
				Message:    err.Error(),
				StatusCode: http.StatusInternalServerError,
			}
		}
		return nil, &githubErrResp
	}

	var user github.GetUserResponse
	err = json.Unmarshal(respBodyJSON, &user)
	if err != nil {
		return nil, &github.ErrorResponse{
			Message:    err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	return &user, nil
}
