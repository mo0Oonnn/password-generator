package response_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mo0Oonnn/password-generator/internal/lib/api/response"
	"github.com/stretchr/testify/assert"
)

func TestOK(t *testing.T) {
	resp := response.OK()

	assert.Equal(t, "OK", resp.Status)
}

type PasswordResponse struct {
	response.Response
	Password string `json:"password"`
	Size     int    `json:"size"`
}

func TestError(t *testing.T) {
	resp := response.Error("error")

	assert.Equal(t, "Error", resp.Status)
}

func TestPasswordResponseOK(t *testing.T) {
	req, err := http.NewRequest("GET", "/test", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	password := "secure_password"
	size := 10

	response.PasswordResponseOK(rr, req, password, size, false)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"status":"OK","password":"secure_password","size":10,"useSpecialChars":false}
`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
	req, err = http.NewRequest("GET", "/test", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	response.PasswordResponseOK(rr, req, password, size, true)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected = `{"status":"OK","password":"secure_password","size":10,"useSpecialChars":true}
`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
