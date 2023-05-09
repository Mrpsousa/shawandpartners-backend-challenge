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
	expected := &entity.GitHubUserResponse{User: entity.GitHubUser{Login: "octocat", ID: 583231, NodeID: "MDQ6VXNlcjU4MzIzMQ==", AvatarUrl: "https://avatars.githubusercontent.com/u/583231?v=4", GravatarID: "", Url: "https://api.github.com/users/octocat", HtmlUrl: "https://github.com/octocat", FollowersUrl: "https://api.github.com/users/octocat/followers", FollowingUrl: "https://api.github.com/users/octocat/following{/other_user}", GistsUrl: "https://api.github.com/users/octocat/gists{/gist_id}", StarredUrl: "https://api.github.com/users/octocat/starred{/owner}{/repo}", SubscriptionsUrl: "https://api.github.com/users/octocat/subscriptions", OrganizationsUrl: "https://api.github.com/users/octocat/orgs", ReposUrl: "https://api.github.com/users/octocat/repos", EventsUrl: "https://api.github.com/users/octocat/events{/privacy}", ReceivedEventsUrl: "https://api.github.com/users/octocat/received_events", Type: "User", SiteAdmin: false, Name: "The Octocat", Company: "@github", Blog: "https://github.blog", Location: "San Francisco", Email: "", Hireable: "", Bio: "", TwitterUsername: "", PublicRepos: 8, PublicGists: 8, Followers: 9149, Following: 9, CreatedAt: "2011-01-25T18:44:36Z", UpdatedAt: "2023-04-22T11:18:59Z"}}
	model := &entity.GitHubUserResponse{}
	client := http.Client{}
	gituser := service.NewGitUser(client)
	response, err := gituser.GitUserDetails(value)

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.IsType(t, model, response)
	assert.Equal(t, expected, response)
}
