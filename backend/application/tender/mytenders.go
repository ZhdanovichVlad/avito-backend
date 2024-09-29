package tender

import (
	"avitoTest/backend/domain/tender"
	"errors"
	"fmt"
	"net/http"
)

func (s NewTenderApplication) GetMyTenders(limit, offset, username string, tenders *[]tender.Tender) (httpCode int, msgErr error) {
	const op = "application.tender.mytenders."

	limitInt, offsetInt, err := s.DataValidator.LimitAndOffsetCheck(limit, offset)
	if err != nil {
		msgErr = errors.New("connot convert limit or offset to int")
		httpCode = http.StatusBadRequest
		return
	}

	httpCode, err = s.DataValidator.UserNameExist(username)
	if err != nil {
		msgErr = fmt.Errorf("op - %s. Wrong user name. err - %w", err)
		return
	}

	searchType := UsernameSearch
	*tenders = make([]tender.Tender, 0, limitInt)

	err = s.TenderInfrastructure.GetTenders(tenders, limitInt, offsetInt, username, searchType)
	if err != nil {
		err = fmt.Errorf("error when receiving data from the server. %w", err)
		httpCode = http.StatusInternalServerError
		return
	}

	return http.StatusOK, err

}
