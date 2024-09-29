package tender

import (
	"errors"
	"fmt"
	"net/http"

	"avitoTest/backend/domain/tender"
)

func (serever NewTenderApplication) GetTender(limit, offset, serviceType string, tenders *[]tender.Tender) (httpCode int, msgErr error) {
	const op = "backend.application.tender.gettenders"

	limitInt, offsetInt, err := serever.DataValidator.LimitAndOffsetCheck(limit, offset)
	if err != nil {
		msgErr = errors.New("connot convert limit or offset to int")
		httpCode = http.StatusBadRequest
		return
	}

	if serviceType != "" && serever.DataValidator.IsTenderServiceTypeIncorrect(serviceType) {
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

	err = serever.TenderInfrastructure.GetTenders(tenders, limitInt, offsetInt, serviceType, searchType)
	if err != nil {
		err = fmt.Errorf("error when receiving data from the server. %w", err)
		httpCode = http.StatusInternalServerError
		return
	}

	return http.StatusOK, err
}
