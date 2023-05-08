package usecase

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/mrpsousa/api/internal/entity"
)

var URLBase = "https://api.github.com/users"

type GitUser interface {
	GitUserSince(since string) (*entity.UsersReponse, error)
	GitUserDetails(userName string) (*entity.GitHubUserResponse, error)
	GitUserRepos(UserName string) (*entity.RepoReponse, error)
}

type gitUsers struct {
	client http.Client
}

func NewGitUser(client http.Client) GitUser {
	return &gitUsers{
		client: client,
	}
}

func (s *gitUsers) GitUserSince(since string) (*entity.UsersReponse, error) {
	var usersReponse = entity.UsersReponse{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s?since%s=3&per_page=10", URLBase, since), nil)
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
	json.NewDecoder(resp.Body).Decode(&users)

	var userArray = make([]string, 0, len(users))
	for _, user := range users {
		userArray = append(userArray, user.Login)
	}
	usersReponse.Users = userArray

	return &usersReponse, nil
}

func (s *gitUsers) GitUserDetails(userName string) (*entity.GitHubUserResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/octocat%s", URLBase, userName), nil)
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

	var userDetail entity.GitHubUserResponse
	json.NewDecoder(resp.Body).Decode(&userDetail.User)

	return &userDetail, nil
}

func (s *gitUsers) GitUserRepos(userName string) (*entity.RepoReponse, error) {
	var response = entity.RepoReponse{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s/repos", URLBase, userName), nil)
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
	json.NewDecoder(resp.Body).Decode(&userRepos)

	var repoArray = make([]string, 0, len(userRepos))
	for _, repo := range userRepos {
		repoArray = append(repoArray, repo.Name)
	}
	response.Repositories = repoArray
	return &response, nil

}
