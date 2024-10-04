package tender

import (
	"avitoTest/backend/internal/domain/tender"
	"fmt"
	"net/http"
	"time"
)

// NewTwnder function for creating a new tender
func (server *Application) NewTender(tender *tender.Tender) (httpCode int, err error) {
	const op = "application.tender.newtender"
	tender.Status = "Created"
	tender.CreatedAt = time.Now()
	tender.Version = 1

	err = tender.FullValidate()
	if err != nil {
		msgErr := fmt.Errorf("op -%s. Failed to create a new tender %w", op, err)
		return http.StatusBadRequest, msgErr

	}
	httpCode, err = server.DataValidator.UserAndOrgExists(tender.CreatorUsername, tender.OrganizationId)
	if err != nil {
		msgErr := fmt.Errorf("op - %s. Error - %w.", op, err)
		return httpCode, msgErr
	}

	id, err := server.TenderInfrastructure.CreateTender(tender)
	if err != nil {
		msgErr := fmt.Errorf("op - %s. Failed to create a new tender %w", op, err)
		return http.StatusInternalServerError, msgErr
	}
	tender.Id = id

	return http.StatusCreated, err
}
