package web_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	web "github.com/mrpsousa/api/internal/infra/web"
	"github.com/stretchr/testify/assert"
)

func TestUserDetailHandler(t *testing.T) {
	reqUrl := fmt.Sprintf("/api/users/%v/details", "octocat")
	req, err := http.NewRequest("GET", reqUrl, nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(web.UserDetailsHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	assert.NotEmpty(t, rr.Body)
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestSinceHandlerWithOutValueError(t *testing.T) {
	reqUrl := fmt.Sprintf("/api/users/%v/details", "")
	req, err := http.NewRequest("GET", reqUrl, nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(web.UserDetailsHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}
