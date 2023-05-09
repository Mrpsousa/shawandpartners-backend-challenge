package docs

import "net/http"

func GetSwaggerHTML(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./docs/index.html")
}

func GetSwaggerYAML(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./docs/swagger.yaml")
}
