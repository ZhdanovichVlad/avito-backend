package dto

import (
	"avitoTest/backend/internal/entity"
	"time"
)

type TenderHandler struct {
	Id              string    `json:"id,omitempty"`
	Name            string    `json:"name,omitempty"`
	Description     string    `json:"description,omitempty"`
	ServiceType     string    `json:"serviceType,omitempty"`
	Status          string    `json:"status,omitempty"`
	OrganizationId  string    `json:"organizationId,omitempty"`
	CreatorUsername string    `json:"creatorUsername,omitempty"`
	Version         int       `json:"verstion,omitempty"`
	CreatedAt       time.Time `json:"createdAt,omitempty"`
}

func TenderUseCaseToTenderDTO(tender entity.Tender) TenderHandler {
	tenderHandler := TenderHandler{}
	tenderHandler.Id = tender.Id
	tenderHandler.Name = tender.Name
	tenderHandler.Description = tender.Description
	tenderHandler.ServiceType = tender.ServiceType
	tenderHandler.Status = tender.Status
	tenderHandler.OrganizationId = tender.OrganizationId
	tenderHandler.CreatorUsername = tender.CreatorUsername
	tenderHandler.Version = tender.Version
	tenderHandler.CreatedAt = tender.CreatedAt
	return tenderHandler
}
