package ping

import (
	"avitoTest/backend/internal/presentation/http/responce"
	"net/http"
)

func ping(writer http.ResponseWriter, request *http.Request) {
	msg := "ping"
	responce.AnswerSuccess(writer, request, 200, msg)
	return
}
