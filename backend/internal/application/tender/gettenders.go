package tender

import (
	"avitoTest/backend/internal/domain/tender"
	"errors"
	"fmt"
	"net/http"
)

func (server Application) GetTender(limit, offset, serviceType string, tenders *[]tender.Tender) (httpCode int, msgErr error) {
	limitInt, offsetInt, err := server.DataValidator.LimitAndOffsetCheck(limit, offset)
	if err != nil {
		msgErr = errors.New("connot convert limit or offset to int")
		httpCode = http.StatusBadRequest
		return
	}

	if serviceType != "" && server.DataValidator.IsTenderServiceTypeIncorrect(serviceType) {
		msgErr = fmt.Errorf("incorrectly specified tender service type")
		httpCode = http.StatusBadRequest
		return
	}

	var searchType int
	if serviceType == "" {
		searchType = ServiceTypeIsEmpty
	} else {
		searchType = ServiceTypeNotEmpty
	}

	*tenders = make([]tender.Tender, 0, limitInt)

	err = server.TenderInfrastructure.GetTenders(tenders, limitInt, offsetInt, serviceType, searchType)
	if err != nil {
		err = fmt.Errorf("error when receiving data from the server. %w", err)
		httpCode = http.StatusInternalServerError
		return
	}

	return http.StatusOK, err
}
