package tender

import (
	"avitoTest/backend/internal/domain/tender"
	"avitoTest/backend/internal/infrastructure/storageerror"
	"errors"
	"fmt"
	"net/http"
)

// ChencgeTenderStatus Function for changing the status of a tender
func (server Application) ChencgeTenderStatus(tender *tender.Tender, tenderId, status, username string) (httpCode int, errMsg error) {
	const op = "application.tender.changetenderstatus"

	err := server.TenderInfrastructure.GetFullTender(tender, tenderId)
	if err != nil {
		if errors.Is(err, storageerror.ErrTenderNotFound) {
			msgErr := fmt.Errorf("op - %s tender not found. %w", op, err)
			return http.StatusNotFound, msgErr
		}
		msgErr := fmt.Errorf("op - %s. failed to retrieve the tender from the database. %w", op, err)
		return http.StatusInternalServerError, msgErr
	}

	httpCode, err = server.DataValidator.UserAndOrgExists(username, tender.OrganizationId)
	if err != nil {
		msgErr := fmt.Errorf("op - %s. Error - %w.", op, err)
		return httpCode, msgErr
	}

	err = server.TenderInfrastructure.UpdateTenderStatus(tenderId, status)
	if err != nil {
		msgErr := fmt.Errorf("op - %s. Failed to update the status of the tender. Error - %w.", op, err)
		return http.StatusInternalServerError, msgErr
	}

	tender.Status = status
	return http.StatusCreated, nil
}
