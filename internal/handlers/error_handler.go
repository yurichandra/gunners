package handlers

import (
	"net/http"

	"github.com/go-chi/render"
)

type errorResponse struct {
	status  int
	Message string `json:"message"`
}

func newErrorResponse(status int, message string) *errorResponse {
	return &errorResponse{
		status:  status,
		Message: message,
	}
}

func (res *errorResponse) Render(writer http.ResponseWriter, request *http.Request) error {
	render.Status(request, res.status)
	return nil
}
