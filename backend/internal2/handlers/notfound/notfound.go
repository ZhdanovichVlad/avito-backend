package notfound

import (
	"avitoTest/backend/internal2/lib/api/response"
	"errors"
	"net/http"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	const op = "backend.handlers.notFound"
	err := errors.New("invalid URL request")
	response.AnswerError(w, r, op, http.StatusNotFound, err)
	return
}
