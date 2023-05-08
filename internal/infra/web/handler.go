package web

import (
	"encoding/json"
	"fmt"
	http "net/http"

	"github.com/mrpsousa/api/internal/usecase"
)

var (
	client = http.Client{}
)

func SinceHandler(w http.ResponseWriter, r *http.Request) {
	msgDecoded, err := DecodeSinceDateFromURL(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(nil)
		fmt.Println(err) //may be a log
		return
	}
	gituser := usecase.NewGitUser(client)
	output, err := gituser.GitUserSince(*msgDecoded)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(nil)
		fmt.Println(err) //may be a log
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

func UserDetailsHandler(w http.ResponseWriter, r *http.Request) {
	msgDecoded, err := DecodeUserNameToDetailFromURL(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(nil)
		fmt.Println(err) //may be a log
		return
	}

	gituser := usecase.NewGitUser(client)
	output, err := gituser.GitUserDetails(*msgDecoded)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(nil)
		fmt.Println(err) //may be a log
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

func UserReposHandler(w http.ResponseWriter, r *http.Request) {
	msgDecoded, err := DecodeUserNameToReposFromURL(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(nil)
		fmt.Println(err) //may be a log
		return
	}
	client := &http.Client{}
	gituser := usecase.NewGitUser(*client)
	output, err := gituser.GitUserRepos(*msgDecoded)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(nil)
		fmt.Println(err) //may be a log
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}
