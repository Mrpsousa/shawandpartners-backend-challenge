package usecase_test

import (
	"net/http"
	"testing"

	"github.com/mrpsousa/api/internal/entity"
	service "github.com/mrpsousa/api/internal/usecase"
	"github.com/stretchr/testify/assert"
)

func TestSinceSuccess(t *testing.T) {
	value := "2"
	client := http.Client{}
	gituser := service.NewGitUser(client)
	response, err := gituser.GitUserSince(value)

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.IsType(t, entity.UsersReponse{}, *response)
}
