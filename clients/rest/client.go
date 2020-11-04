package rest

import (
	"bytes"
	"encoding/json"
	"net/http"
)

var (
	//Client REST client interface
	Client Interface
)

func init() {
	Client = &client{client: &http.Client{}}
}

// Interface defines the HTTP Methods for a RESTful client
type Interface interface {
	GET(endpoint string, headers *http.Header) (*http.Response, error)
	POST(endpoint string, headers *http.Header, body interface{}) (*http.Response, error)
	PUT(endpoint string, headers *http.Header, body interface{}) (*http.Response, error)
	// 	DELETE(endpoint string) (*http.Response, error)
	// 	PATCH(endpoint string, body interface{}) (*http.Response, error)
	// 	HEAD(endpoint string) (*http.Response, error)
}

type client struct {
	client *http.Client
}

func (c *client) GET(endpoint string, headers *http.Header) (*http.Response, error) {

	request, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}
	request.Header = *headers

	return c.client.Do(request)
}

func (c *client) POST(endpoint string, headers *http.Header, body interface{}) (*http.Response, error) {
	jsonBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewReader(jsonBytes))
	if err != nil {
		return nil, err
	}
	request.Header = *headers

	return c.client.Do(request)
}

func (c *client) PUT(endpoint string, headers *http.Header, body interface{}) (*http.Response, error) {
	jsonBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(http.MethodPut, endpoint, bytes.NewReader(jsonBytes))
	if err != nil {
		return nil, err
	}
	request.Header = *headers

	return c.client.Do(request)
}
