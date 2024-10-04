package tender

import (
	tenderApplication "avitoTest/backend/internal/application/tender"
	tender2 "avitoTest/backend/internal/domain/tender"
	"avitoTest/backend/internal/presentation/http/responce"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
)

func PutTenderStatus(s tenderApplication.Application) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		const op = "presentation.http.tenders.puttenderstatus"

		tenderId := chi.URLParam(request, "tenderId")
		reqQuery := request.URL.Query()
		status := reqQuery.Get("status")
		username := reqQuery.Get("username")

		var tender = tender2.Tender{}
		httpCode, err := s.ChencgeTenderStatus(&tender, tenderId, status, username)
		if err != nil {
			responce.AnswerError(writer, request, op, httpCode, err)
			return
		}
		render.Status(request, httpCode)
		render.JSON(writer, request, tender)
	}
}
