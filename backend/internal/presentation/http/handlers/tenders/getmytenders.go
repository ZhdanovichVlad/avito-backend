package tender

import (
	tenderApplication "avitoTest/backend/internal/application/tender"
	"avitoTest/backend/internal/domain/tender"
	"avitoTest/backend/internal/presentation/http/responce"
	"github.com/go-chi/render"
	"net/http"
)

// GetMyTendersH function for obtaining the list of user's tenders
func GetMyTendersH(s tenderApplication.Application) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		const op = "interfacsec.http.hadnlers.tenders.getmytendred"

		reqQuery := request.URL.Query()
		limit := reqQuery.Get("limit")
		offset := reqQuery.Get("offset")
		usernsme := reqQuery.Get("username")

		var tenders []tender.Tender
		httpCode, err := s.GetMyTenders(limit, offset, usernsme, &tenders)
		if err != nil {
			responce.AnswerError(writer, request, op, httpCode, err)
			return
		}
		render.Status(request, httpCode)
		render.JSON(writer, request, tenders)
	}
}
