package usecase_test

import (
	"net/http"
	"testing"

	service "github.com/mrpsousa/api/internal/usecase"
	"github.com/stretchr/testify/assert"
)

func TestReposSuccess(t *testing.T) {
	value := "octocat"
	expected := []string([]string{"mojombo", "defunkt", "pjhyett", "wycats", "ezmobius", "ivey", "evanphx", "vanpelt", "wayneeseguin", "brynary"})
	client := http.Client{}
	gituser := service.NewGitUser(client)
	response, err := gituser.GitUserSince(value)

	assert.Nil(t, err)
	assert.NotNil(t, response)
	var array = make([]string, 0, len(response))
	assert.IsType(t, array, response)
	assert.Equal(t, response, expected)

}
