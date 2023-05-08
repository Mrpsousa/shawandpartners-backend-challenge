package usecase

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/mrpsousa/api/internal/entity"
)

type GitUser interface {
	GitUserSince(since string) ([]string, error)
	GitUserDetails(userName string) (*entity.GitHubUser, error)
	GitUserRepos(UserName string) ([]string, error)
}

type gitUsers struct {
	client http.Client
}

func NewGitUser(client http.Client) GitUser {
	return &gitUsers{
		client: client,
	}
}

func (s *gitUsers) GitUserSince(since string) ([]string, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.github.com/users?since%s=3&per_page=10", since), nil)
	if err != nil {
		err := errors.New("creating_request: get_users_since")
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		err := errors.New("sending_request: get_users_since")
		return nil, err
	}

	defer resp.Body.Close()

	var users []entity.UserSince

	err = json.NewDecoder(resp.Body).Decode(&users)
	if err != nil {
		panic(err)
	}

	var userArray = make([]string, 0, len(users))
	for _, user := range users {
		userArray = append(userArray, user.Login)
	}

	return userArray, nil
}

func (s *gitUsers) GitUserDetails(userName string) (*entity.GitHubUser, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.github.com/users/octocat%s", userName), nil)
	if err != nil {
		err := errors.New("creating_request: get_users_details")
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		err := errors.New("sending_request: get_users_details")
		return nil, err
	}
	defer resp.Body.Close()

	var userDetail entity.GitHubUser
	err = json.NewDecoder(resp.Body).Decode(&userDetail)
	if err != nil {
		err := errors.New("deconding_json")
		return nil, err
	}

	return &userDetail, nil
}

func (s *gitUsers) GitUserRepos(userName string) ([]string, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.github.com/users/%s/repos", userName), nil)
	if err != nil {
		err := errors.New("creating_request: get_users_details")
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		err := errors.New("sending_request: get_users_details")
		return nil, err
	}
	defer resp.Body.Close()

	var userRepos []entity.Repository
	err = json.NewDecoder(resp.Body).Decode(&userRepos)
	if err != nil {
		panic(err)
	}
	var repoArray = make([]string, 0, len(userRepos))
	for _, repo := range userRepos {
		repoArray = append(repoArray, repo.Name)
	}

	return repoArray, nil

}
