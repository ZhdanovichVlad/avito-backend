package tender

import (
	"avitoTest/backend/domain/tender"
	"avitoTest/backend/infrastructure/storageerror"
	"errors"
	"fmt"
	"net/http"
)

func (server NewTenderApplication) GetTenderStatus(tenderId, username string, status *string) (httpCode int, err error) {
	const op = "application.tender.tenderStatus"

	var tender tender.Tender

	err = server.TenderInfrastructure.GetFullTender(&tender, tenderId)
	if err != nil {
		if errors.Is(err, storageerror.ErrTenderNotFound) {
			msgErr := fmt.Errorf("op - %s tender not found. %w", op, err)
			return http.StatusNotFound, msgErr
		}
		msgErr := fmt.Errorf("op - %s. failed to retrieve the tender from the database. %w", op, err)
		return http.StatusInternalServerError, msgErr
	}

	httpCode, err = server.DataValidator.FullExist(username, tender.OrganizationId)
	if err != nil {
		msgErr := fmt.Errorf("op - %s. Error - %w.", op, err)
		return httpCode, msgErr
	}

	*status = tender.Status
	return httpCode, nil
}
