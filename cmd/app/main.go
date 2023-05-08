package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mrpsousa/api/internal/infra/web"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/users", web.SinceHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/users/{username}/details", web.UserDetailsHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/users/{username}/repos", web.UserReposHandler).Methods(http.MethodGet)
	log.Println("Starting server on port 8080...")

	http.ListenAndServe(":8080", r)

}
