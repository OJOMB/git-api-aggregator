package githubprovider

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/OJOMB/git-api-aggregator/clients/rest"
	"github.com/OJOMB/git-api-aggregator/domain/github"
	"github.com/stretchr/testify/assert"
)

func init() {
	rest.Client = &mockClient{}
}

var post func(endpoint string, headers *http.Header, body interface{}) (*http.Response, error)

type errorWhenRead struct{}

func (r errorWhenRead) Read(p []byte) (n int, err error) {
	return 0, errors.New("Terrible Error")
}

func (r errorWhenRead) Close() error {
	return nil
}

type mockClient struct {
	rest.Interface
}

func (mc *mockClient) POST(endpoint string, headers *http.Header, body interface{}) (*http.Response, error) {
	return post(endpoint, headers, body)
}

func TestGetAuthorizationHeader(t *testing.T) {
	expected := "token 000"
	result := GetAuthorizationHeader("000")

	assert.Equal(t, expected, result)
}

func TestCreateRepo(t *testing.T) {
	testTable := []struct {
		Number       int
		PostFunction func(endpoint string, headers *http.Header, body interface{}) (*http.Response, error)
		Request      *github.CreateRepoRequest
		UserName     string
		Token        string
		ExpectedResp *github.CreateRepoResponse
		ExpectedErr  *github.ErrorResponse
	}{
		{
			// case when post request to github returns error
			Number: 0,
			PostFunction: func(endpoint string, headers *http.Header, body interface{}) (*http.Response, error) {
				return nil, errors.New("Request attempt failed")
			},
			Request:      &github.CreateRepoRequest{},
			UserName:     "testuser",
			Token:        "testtoken",
			ExpectedResp: nil,
			ExpectedErr: &github.ErrorResponse{
				StatusCode: http.StatusInternalServerError,
				Message:    "Request attempt failed",
			},
		},
		{
			// case when response body returned is unreadable
			Number: 1,
			PostFunction: func(endpoint string, headers *http.Header, body interface{}) (*http.Response, error) {
				return &http.Response{Body: errorWhenRead{}}, nil
			},
			Request:      &github.CreateRepoRequest{},
			UserName:     "testuser",
			Token:        "testtoken",
			ExpectedResp: nil,
			ExpectedErr: &github.ErrorResponse{
				StatusCode: http.StatusInternalServerError,
				Message:    "Terrible Error",
			},
		},
		{
			// case when github error returned by provider
			Number: 2,
			PostFunction: func(endpoint string, headers *http.Header, body interface{}) (*http.Response, error) {
				return &http.Response{
					Body:       ioutil.NopCloser(strings.NewReader(`{"Message": "test1", "Documentation_url":"test2", "errors":[{"resource":"12345","code":"500","message":"test3","field":"test4"}]}`)),
					StatusCode: http.StatusInternalServerError,
				}, nil
			},
			Request:      &github.CreateRepoRequest{},
			UserName:     "testuser",
			Token:        "testtoken",
			ExpectedResp: nil,
			ExpectedErr: &github.ErrorResponse{
				StatusCode:       http.StatusInternalServerError,
				Message:          "test1",
				DocumentationURL: "test2",
				Errors: []github.Error{
					{Resource: "12345", Code: "500", Message: "test3", Field: "test4"},
				},
			},
		},
		{
			// case when create repo request receives successful response
			Number: 3,
			PostFunction: func(endpoint string, headers *http.Header, body interface{}) (*http.Response, error) {
				return &http.Response{
					Body:       ioutil.NopCloser(strings.NewReader(`{"id": 1296269,"name": "Hello-Test","full_name": "octocat/Hello-World","owner": {"login": "octocat","id": 1,"html_url": "HTMLURLTest","url": "urlTest"},"permission": "permissionTest"}`)),
					StatusCode: http.StatusCreated,
				}, nil
			},
			Request:  &github.CreateRepoRequest{},
			UserName: "testuser",
			Token:    "testtoken",
			ExpectedResp: &github.CreateRepoResponse{
				ID:         1296269,
				Name:       "Hello-Test",
				FullName:   "octocat/Hello-World",
				Permission: "permissionTest",
				Owner: github.Owner{
					ID:      1,
					Login:   "octocat",
					HTMLURL: "HTMLURLTest",
					URL:     "urlTest",
				},
			},
			ExpectedErr: nil,
		},
		{
			// case when response has unexpected status code
			Number: 4,
			PostFunction: func(endpoint string, headers *http.Header, body interface{}) (*http.Response, error) {
				return &http.Response{
					Body:       ioutil.NopCloser(strings.NewReader(`{"id": 1296269}`)),
					StatusCode: 100,
				}, nil
			},
			Request:      &github.CreateRepoRequest{},
			UserName:     "testuser",
			Token:        "testtoken",
			ExpectedResp: nil,
			ExpectedErr: &github.ErrorResponse{
				StatusCode: http.StatusInternalServerError,
				Message:    "Received response with an unexpected status code: 100",
			},
		},
		{
			// case when github response contains bad JSON
			Number: 5,
			PostFunction: func(endpoint string, headers *http.Header, body interface{}) (*http.Response, error) {
				return &http.Response{
					Body:       ioutil.NopCloser(strings.NewReader(`{sionTest"}`)),
					StatusCode: http.StatusCreated,
				}, nil
			},
			Request:      &github.CreateRepoRequest{},
			UserName:     "testuser",
			Token:        "testtoken",
			ExpectedResp: nil,
			ExpectedErr: &github.ErrorResponse{
				StatusCode: http.StatusInternalServerError,
				Message:    "Received invalid JSON response from provider: invalid character 's' looking for beginning of object key string",
			},
		},
		{
			// case when github returns error response with bad JSON
			Number: 5,
			PostFunction: func(endpoint string, headers *http.Header, body interface{}) (*http.Response, error) {
				return &http.Response{
					Body:       ioutil.NopCloser(strings.NewReader(`{sionTest"}`)),
					StatusCode: http.StatusInternalServerError,
				}, nil
			},
			Request:      &github.CreateRepoRequest{},
			UserName:     "testuser",
			Token:        "testtoken",
			ExpectedResp: nil,
			ExpectedErr: &github.ErrorResponse{
				StatusCode: http.StatusInternalServerError,
				Message:    "Received invalid JSON response from provider: invalid character 's' looking for beginning of object key string",
			},
		},
	}

	for _, test := range testTable {
		t.Run(
			fmt.Sprintf("CreateRepo test %d", test.Number),
			func(t *testing.T) {
				post = test.PostFunction
				resultResp, resultErr := CreateRepo(test.Request, test.UserName, test.Token)

				if !reflect.DeepEqual(resultResp, test.ExpectedResp) {
					t.Errorf(
						"Unexpected response value\nExpected:\n    %v\nGot:\n    %v",
						test.ExpectedResp, resultResp,
					)
				}
				if !reflect.DeepEqual(resultErr, test.ExpectedErr) {
					t.Errorf(
						"Unexpected error value\nExpected:\n    %v\nGot:\n    %v",
						test.ExpectedErr, resultErr,
					)
				}
			},
		)
	}
}
