package github

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateRepoBodyToJSON(t *testing.T) {
	// gr := &githubRepos{}

	testBody := CreateRepoRequest{
		Name:        "golang introduction",
		Description: "a golang introduction repo",
		Homepage:    "https://github.com",
		Private:     true,
		HasIssues:   false,
		HasProjects: true,
		HasWiki:     false,
	}

	expected := []byte(`{"name":"golang introduction","description":"a golang introduction repo","homepage":"https://github.com","private":true,"has_issues":false,"has_projects":true,"has_wiki":false}`)

	result, err := json.Marshal(&testBody)
	assert.NoError(t, err)
	assert.EqualValues(t, result, expected)
}

func TestJSONToCreateRepoBody(t *testing.T) {
	// gr := &githubRepos{}

	expected := CreateRepoRequest{
		Name:        "golang introduction",
		Description: "a golang introduction repo",
		Homepage:    "https://github.com",
		Private:     true,
		HasIssues:   false,
		HasProjects: true,
		HasWiki:     false,
	}

	var result CreateRepoRequest

	jsonBytes := []byte(`{"name":"golang introduction","description":"a golang introduction repo","homepage":"https://github.com","private":true,"has_issues":false,"has_projects":true,"has_wiki":false}`)

	err := json.Unmarshal(jsonBytes, &result)
	assert.NoError(t, err)

	assert.EqualValues(t, result, expected)
}
