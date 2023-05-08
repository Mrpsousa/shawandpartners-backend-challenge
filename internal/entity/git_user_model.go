package entity

type UserSince struct {
	Login string `json:"login"`
}

type GitHubUserResponse struct {
	User GitHubUser `json:"user"`
}
type GitHubUser struct {
	Login string `json:"login"`
	Name  string `json:"name"`
	URL   string `json:"url"`
}

type Repository struct {
	Name string `json:"name"`
}

type RepoReponse struct {
	Repositories []string `json:"repositories"`
}

type UsersReponse struct {
	Users []string `json:"users"`
}
