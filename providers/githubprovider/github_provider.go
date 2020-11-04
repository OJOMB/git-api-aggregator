package githubprovider

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/OJOMB/git-api-aggregator/clients/rest"
	"github.com/OJOMB/git-api-aggregator/domain/github"
)

const (
	token                     = ""
	headerAuthorization       = "Authorization"
	headerAuthorizationFormat = "token %s"

	apiCreateRepoFormat = "https://api.github.com/user/%s/%s"
)

func getAuthorizationHeader(token string) string {
	// Header "Authorization": "token 0000000000000000000000000000000000000000"
	return fmt.Sprintf(headerAuthorizationFormat, token)
}

// CreateRepo creates a Github repository.
func CreateRepo(req *github.CreateRepoRequest, userName, token string) (*github.CreateRepoResponse, *github.ErrorResponse) {
	// construct request
	headers := &http.Header{}
	headers.Set("Authorization", getAuthorizationHeader(token))
	endpoint := fmt.Sprintf(apiCreateRepoFormat, userName, req.Name)
	resp, err := rest.Client.POST(endpoint, headers, req)
	if err != nil {
		return nil, &github.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	// read response body
	jsonBytes, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		return nil, &github.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	if resp.StatusCode == http.StatusCreated {
		log.Printf("Received 201 in response to CreateRepo request from provider")
		var createResponse github.CreateRepoResponse
		err = json.Unmarshal(jsonBytes, &createResponse)
		if err != nil {
			return nil, &github.ErrorResponse{
				StatusCode: http.StatusInternalServerError,
				Message:    fmt.Sprintf("Received invalid JSON response from provider: %s", err.Error()),
			}
		}
		return &createResponse, nil
	} else if resp.StatusCode >= 400 && resp.StatusCode <= 599 {
		log.Printf("Received error response to CreateRepo request from provider. Response status code: %d", resp.StatusCode)
		var githubErr github.ErrorResponse
		err = json.Unmarshal(jsonBytes, &githubErr)
		if err != nil {
			return nil, &github.ErrorResponse{
				StatusCode: http.StatusInternalServerError,
				Message:    fmt.Sprintf("Received invalid JSON response from provider: %s", err.Error()),
			}
		}
		githubErr.StatusCode = resp.StatusCode
		return nil, &githubErr
	}

	log.Printf("Received response with an unexpected status code: %d", resp.StatusCode)

	return nil, &github.ErrorResponse{
		StatusCode: http.StatusInternalServerError,
		Message:    fmt.Sprintf("Received response with an unexpected status code: %d", resp.StatusCode),
	}
}
