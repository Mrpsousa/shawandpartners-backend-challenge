package entity

type User struct {
	ID        int    `json:"id"`
	Login     string `json:"login"`
	AvatarURL string `json:"avatar_url"`
}

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
