package tenders

import (
	"avitoTest/backend/internal/lib/models"
	"errors"
	"fmt"
	"net/http"

	"avitoTest/backend/internal/handlers/get"
	"avitoTest/backend/internal/lib/api/limitandoffsetcheck"
	"avitoTest/backend/internal/lib/api/response"
	"avitoTest/backend/internal/lib/api/typecheck"
	"github.com/go-chi/render"
)

// description of service type
const construction = "Construction"
const delivery = "Delivery"
const Manufacture = "Manufacture"

// Interface for working with the database

// GetTenderH Handler to retrieve a list of tenders according to specified conditions. The handle supports pagination
func GetTenderH(server get.ServerGet) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		const op = "backend.internal.handlers.get.GetTenderH"

		reqQuery := request.URL.Query()
		limit := reqQuery.Get("limit")
		offset := reqQuery.Get("offset")
		serviceType := reqQuery.Get("service_type")

		limitInt, offsetInt, err := limitandoffsetcheck.LimitAndOffsetCheck(limit, offset)
		if err != nil {
			err = errors.New("connot convert limit or offset to int")
			response.AnswerError(writer, request, op, http.StatusBadRequest, err)
			return
		}
		if serviceType != "" && typecheck.IsTenderServiceTypeIncorrect(serviceType) {
			msgErr := fmt.Errorf("incorrectly specified tender service type")
			response.AnswerError(writer, request, op, http.StatusBadRequest, msgErr)
			return
		}

		var searchType int
		if serviceType == "" {
			searchType = get.ServiceTypeIsEmpty
		} else {
			searchType = get.ServiceTypeNotEmpty
		}
		tenders := make([]models.Tender, 0, limitInt)
		err = server.GetTenders(&tenders, limitInt, offsetInt, serviceType, searchType)
		if err != nil {
			err = fmt.Errorf("error when receiving data from the server. %s", err)
			response.AnswerError(writer, request, op, http.StatusInternalServerError, fmt.Errorf("failed to retrieve data from the database %w", err))
			return
		}

		render.Status(request, http.StatusOK)
		render.JSON(writer, request, tenders)
	}
}
