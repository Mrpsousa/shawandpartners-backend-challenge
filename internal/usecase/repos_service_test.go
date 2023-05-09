package usecase_test

import (
	"net/http"
	"testing"

	"github.com/mrpsousa/api/internal/entity"
	service "github.com/mrpsousa/api/internal/usecase"
	"github.com/stretchr/testify/assert"
)

func TestReposSuccess(t *testing.T) {
	value := "octocat"
	expected := entity.RepoReponse(entity.RepoReponse{Repositories: []string{"boysenberry-repo-1", "git-consortium", "hello-worId", "Hello-World", "linguist", "octocat.github.io", "Spoon-Knife", "test-repo1"}})
	client := http.Client{}
	gituser := service.NewGitUser(client)
	response, err := gituser.GitUserRepos(value)

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.IsType(t, entity.RepoReponse{}, *response)
	assert.Equal(t, expected, *response)

}
