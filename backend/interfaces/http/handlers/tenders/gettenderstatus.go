package tenders

import (
	"avitoTest/backend/application/tender"
	"avitoTest/backend/interfaces/http/responce"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
)

func GetTenderStatusH(s tender.NewTenderApplication) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		const op = "interfaces.http.handlers.tenders.gettende"
		tenderId := chi.URLParam(request, "tenderId")

		username := request.URL.Query().Get("username")

		var status string
		httpcode, err := s.GetTenderStatus(tenderId, username, &status)
		if err != nil {
			responce.AnswerError(writer, request, op, httpcode, err)
			return
		}

		render.Status(request, httpcode)
		render.JSON(writer, request, status)
	}
}
