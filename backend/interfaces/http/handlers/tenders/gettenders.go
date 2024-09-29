package tenders

import (
	"net/http"

	tenderApplication "avitoTest/backend/application/tender"
	"avitoTest/backend/domain/tender"
	"avitoTest/backend/interfaces/http/responce"

	"github.com/go-chi/render"
)

// GetTendersH Handler to retrieve a list of tenders according to specified conditions. The handle supports pagination
func GetTendersH(s tenderApplication.NewTenderApplication) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		const op = "backend.internal.handlers.get.GetTenderH"

		reqQuery := request.URL.Query()
		limit := reqQuery.Get("limit")
		offset := reqQuery.Get("offset")
		serviceType := reqQuery.Get("service_type")

		var tenders []tender.Tender

		httpCode, err := s.GetTender(limit, offset, serviceType, &tenders)
		if err != nil {
			responce.AnswerError(writer, request, op, httpCode, err)
		}
		render.Status(request, httpCode)
		render.JSON(writer, request, tenders)
	}
}
