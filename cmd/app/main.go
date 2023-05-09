package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mrpsousa/api/docs"
	"github.com/mrpsousa/api/internal/infra/web"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/users", web.SinceHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/users/{username}/details", web.UserDetailsHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/users/{username}/repos", web.UserReposHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/docs", docs.GetSwaggerHTML).Methods(http.MethodGet)
	r.HandleFunc("/api/swagger.yaml", docs.GetSwaggerYAML).Methods("GET")
	r.HandleFunc("/api/coverage", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./coverage.html")
	})

	log.Println("Starting server on port 8080...")

	http.ListenAndServe(":8080", r)

}
