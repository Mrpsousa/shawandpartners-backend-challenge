package web_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	web "github.com/mrpsousa/api/internal/infra/web"
	"github.com/stretchr/testify/assert"
)

func TestUserReposHandler(t *testing.T) {
	reqUrl := fmt.Sprintf("/api/users/%v/repos", "octocat")
	req, err := http.NewRequest("GET", reqUrl, nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(web.UserReposHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	assert.NotEmpty(t, rr.Body)
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestUserReposHandlerWithOutValueError(t *testing.T) {
	reqUrl := fmt.Sprintf("/api/users/%v/repos", "")
	req, err := http.NewRequest("GET", reqUrl, nil)
	expected := string("null\n")
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(web.UserReposHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Equal(t, expected, rr.Body.String())
}

func TestUserReposHandlerWrongValueError(t *testing.T) {
	reqUrl := fmt.Sprintf("/api/users/%v/repos", "89425uup89354897tyu3hgf8349uytrtru3489ru34")
	req, err := http.NewRequest("GET", reqUrl, nil)
	expected := string("{\"repositories\":[]}\n")
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(web.UserReposHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, expected, rr.Body.String())
}
