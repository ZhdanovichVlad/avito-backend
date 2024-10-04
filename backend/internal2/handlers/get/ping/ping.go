package ping

import (
	"avitoTest/backend/internal2/lib/api/response"
	"net/http"
)

func Ping(writer http.ResponseWriter, request *http.Request) {
	response.AnswerSuccess(writer, request, http.StatusOK, "PING CALLED")
	return
}
