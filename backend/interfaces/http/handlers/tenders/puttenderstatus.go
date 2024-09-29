package tenders

import (
	tenderApplication "avitoTest/backend/application/tender"
	tender2 "avitoTest/backend/domain/tender"
	"avitoTest/backend/interfaces/http/responce"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
)

func PutTenderStatus(s tenderApplication.NewTenderApplication) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		const op = "interfaces.http.tenders.puttenderstatus"

		tenderid := chi.URLParam(request, "tenderId")
		reqQuery := request.URL.Query()
		status := reqQuery.Get("status")
		username := reqQuery.Get("username")

		var tender = tender2.Tender{}
		httpCode, err := s.ChencgeTenderStatus(&tender, tenderid, status, username)
		if err != nil {
			responce.AnswerError(writer, request, op, httpCode, err)
			return
		}
		render.Status(request, httpCode)
		render.JSON(writer, request, tender)
	}
}
