package entity

type UserSince struct {
	Login string `json:"login"`
}

type GitHubUser struct {
	Login string `json:"login"`
	Name  string `json:"name"`
	URL   string `json:"url"`
}

type Repository struct {
	Name string `json:"name"`
}
