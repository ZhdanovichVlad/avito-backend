package tender

import (
	tenderApplication "avitoTest/backend/internal/application/tender"
	"avitoTest/backend/internal/domain/tender"
	"avitoTest/backend/internal/presentation/http/responce"
	"net/http"

	"github.com/go-chi/render"
)

// GetTendersH Handler to retrieve a list of tenders according to specified conditions. The handle supports pagination
func GetTendersH(s tenderApplication.Application) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		const op = "backend.internal2.handlers.get.GetTenderH"

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
