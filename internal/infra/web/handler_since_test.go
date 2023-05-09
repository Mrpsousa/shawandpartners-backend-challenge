package web_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	web "github.com/mrpsousa/api/internal/infra/web"
	"github.com/stretchr/testify/assert"
)

func TestSinceHandler(t *testing.T) {
	reqUrl := fmt.Sprintf("api/users?since=%v", 12)
	req, err := http.NewRequest("GET", reqUrl, nil)
	expected := "{\"users\":[\"mojombo\",\"defunkt\",\"pjhyett\",\"wycats\",\"ezmobius\",\"ivey\",\"evanphx\",\"vanpelt\",\"wayneeseguin\",\"brynary\"]}\n"
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(web.SinceHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	assert.NotEmpty(t, rr.Body)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, expected, rr.Body.String())

}

func TestSinceHandlerNotNumericError(t *testing.T) {
	reqUrl := fmt.Sprintf("api/users?since=%v", "aBc")
	req, err := http.NewRequest("GET", reqUrl, nil)
	expected := string("null\n")

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(web.SinceHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Equal(t, expected, rr.Body.String())

}
