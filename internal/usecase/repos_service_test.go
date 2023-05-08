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
	expected := entity.UsersReponse(entity.UsersReponse{Users: []string{"mojombo", "defunkt", "pjhyett", "wycats", "ezmobius", "ivey", "evanphx", "vanpelt", "wayneeseguin", "brynary"}})
	client := http.Client{}
	gituser := service.NewGitUser(client)
	response, err := gituser.GitUserSince(value)

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.IsType(t, entity.UsersReponse{}, *response)
	assert.Equal(t, expected, *response)

}
