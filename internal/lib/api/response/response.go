package response

import (
	"net/http"

	"github.com/go-chi/render"
)

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

type PasswordResponse struct {
	Response
	Password string `json:"password"`
	Size     int    `json:"size"`
}

const (
	StatusOK    = "OK"
	StatusError = "Error"
)

func OK() Response {
	return Response{
		Status: StatusOK,
	}
}

func Error(msg string) Response {
	return Response{
		Status: StatusError,
		Error:  msg,
	}
}

func PasswordResponseOK(w http.ResponseWriter, r *http.Request, password string, size int) {
	render.JSON(w, r, PasswordResponse{
		Response: OK(),
		Password: password,
		Size:     size,
	})
}
