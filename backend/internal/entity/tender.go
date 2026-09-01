package entity

import (
	"avitoTest/backend/internal/entityjson"
	"avitoTest/backend/pkg/errorsx"
	"time"
)

const (
	TenderServiceTypeConstruction = "Construction"
	TenderServiceTypeDelivery     = "Delivery"
	TenderServiceTypeManufacture  = "Manufacture"
	TenderStatusCreated           = "Created"
	TenderStatusPublished         = "Published"
	TenderStatusClosed            = "Closed"
)

// Tender Structure describing server attributes with tenders
type Tender struct {
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

// NewTender Create new tender entity
func NewTender(tenderJSON entityjson.Tender) *Tender {
	tender := Tender{}
	tender.Name = tenderJSON.Name
	tender.Description = tenderJSON.Description
	tender.ServiceType = tenderJSON.ServiceType
	tender.Status = TenderStatusCreated
	tender.OrganizationId = tenderJSON.OrganizationId
	tender.CreatorUsername = tenderJSON.CreatorUsername
	tender.Version = 1
	return &tender
}

// NewTender Create new tender entity
func NewTenderForUpdate(tenderJSON entityjson.Tender) *Tender {
	tender := Tender{}
	tender.Name = tenderJSON.Name
	tender.Description = tenderJSON.Description
	tender.ServiceType = tenderJSON.ServiceType
	return &tender
}

// ValidationTenderServiceType checks if the Tender type corresponds to the specified values
func ValidationTenderServiceType(serviceType string) error {
	const op = "internal.entity.ValidationTenderServiceType"
	switch serviceType {
	case TenderServiceTypeConstruction:
		return nil
	case TenderServiceTypeDelivery:
		return nil
	case TenderServiceTypeManufacture:
		return nil

	default:
		return errorsx.New(errorsx.ErrBadRequest, "incorrect service type", op, nil)
	}
}

// ValidationTenderStatus checks if the tender status corresponds to the specified values
func ValidationTenderStatus(tenderStatus string) error {
	const op = "internal.entity.ValidationTenderStatus"
	switch tenderStatus {
	case TenderStatusCreated:
		return nil
	case TenderStatusPublished:
		return nil
	case TenderStatusClosed:
		return nil
	default:
		return errorsx.New(errorsx.ErrBadRequest, "incorrect tender status", op, nil)
	}
}
