package usecase_test

import (
	"net/http"
	"testing"

	"github.com/mrpsousa/api/internal/entity"
	service "github.com/mrpsousa/api/internal/usecase"
	"github.com/stretchr/testify/assert"
)

func TestUserDetailSuccess(t *testing.T) {
	value := "octocat"
	expected := &entity.GitHubUserResponse{entity.GitHubUser{Login: "OctocatOctocat", Name: "Christian Amoros", URL: "https://api.github.com/users/OctocatOctocat"}}
	model := &entity.GitHubUserResponse{}
	client := http.Client{}
	gituser := service.NewGitUser(client)
	response, err := gituser.GitUserDetails(value)

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.IsType(t, model, response)
	assert.Equal(t, expected, response)
}
