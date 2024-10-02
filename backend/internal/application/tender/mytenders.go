package tender

import (
	"avitoTest/backend/internal/domain/tender"
	"errors"
	"fmt"
	"net/http"
)

type Tender struct {
	status string
}

func (t *Tender) Publish() {
}

func (server Application) GetMyTenders(limit, offset, username string, tenders *[]tender.Tender) (httpCode int, msgErr error) {
	const op = "application.tender.mytenders."

	limitInt, offsetInt, err := server.DataValidator.LimitAndOffsetCheck(limit, offset)
	if err != nil {
		msgErr = errors.New("connot convert limit or offset to int")
		httpCode = http.StatusBadRequest
		return
	}

	httpCode, err = server.DataValidator.UserNameExist(username)
	if err != nil {
		msgErr = fmt.Errorf("op - %server. Wrong user name. err - %w", err)
		return
	}

	searchType := UsernameSearch
	*tenders = make([]tender.Tender, 0, limitInt)

	err = server.TenderInfrastructure.GetTenders(tenders, limitInt, offsetInt, username, searchType)
	if err != nil {
		err = fmt.Errorf("error when receiving data from the server. %w", err)
		httpCode = http.StatusInternalServerError
		return
	}

	return http.StatusOK, err

}
