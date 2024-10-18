package tender

import (
	"context"
	"fmt"

	tenderEntity "avitoTest/backend/internal/entity/tender"
	"avitoTest/backend/pkg/errorsx"
)

// Create method for creating a new tender
func (server *TenderApplacation) Create(contex context.Context, tender *tenderEntity.Tender) error {
	const op = "usecase.tender.create"
	tender.Status = "Created"

	err := tenderEntity.ValidateTender(tender)
	if err != nil {
		return errorsx.ErrInvalidData
	}

	exist, err := server.tenderRepository.ValidateResponsibleEmployee(tender.OrganizationId, tender.CreatorUsername)
	if err != nil {
		fmt.Println("Ошибка тут 1")
		//msgErr := fmt.Errorf("op - %s. Error - %w.", op, err)
		return errorsx.ErrInternalRepository
	}
	if !exist {
		return errorsx.ErrEmployeeNotResponsible
	}

	id, err := server.tenderRepository.CreateTender(contex, tender)
	if err != nil {
		fmt.Println("Ошибка тут 2")
		//msgErr := fmt.Errorf("op - %s. Failed to create a new tender %w", op, err)
		return errorsx.ErrInternalRepository
	}

	tender.Id = id
	return nil
}
