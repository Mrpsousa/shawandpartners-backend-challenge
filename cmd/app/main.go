package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mrpsousa/api/internal/infra/web"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/users", web.SinceHandler)
	r.HandleFunc("/api/users/{username}/details", web.UserDetailsHandler)
	r.HandleFunc("/api/users/{username}/repos", web.UserReposHandler)
	http.ListenAndServe(":8080", r)

}
