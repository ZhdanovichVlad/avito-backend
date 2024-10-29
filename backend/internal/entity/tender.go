package entity

import (
	"avitoTest/backend/dto"
	"fmt"
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
	Id              string
	Name            string
	Description     string
	ServiceType     string
	Status          string
	OrganizationId  string
	CreatorUsername string
	Version         int
	CreatedAt       time.Time
}

func NewTenderUseCase(tenderDTO dto.TenderHandler) (Tender, error) {
	tender := Tender{}
	tender.Id = tenderDTO.Id
	tender.Name = tenderDTO.Name
	tender.Description = tenderDTO.Description
	tender.Status = TenderStatusCreated
	tender.OrganizationId = tenderDTO.OrganizationId
	tender.CreatorUsername = tenderDTO.CreatorUsername
	tender.Version = 1
	tender.CreatedAt = time.Now()

	if len([]rune(tender.Name)) > 100 && len([]rune(tender.Name)) != 0 {
		return Tender{}, fmt.Errorf("name len more than 100 or emty")
	}
	if len([]rune(tender.Description)) > 500 && len([]rune(tender.Description)) != 0 {
		return Tender{}, fmt.Errorf("description len more than 500 or emty")
	}
	if len([]rune(tender.CreatorUsername)) > 100 && len([]rune(tender.CreatorUsername)) != 0 {
		return Tender{}, fmt.Errorf("description len more than 500 or emty")
	}
	err := ValidationTenderServiceType(tender.ServiceType)
	if err != nil {
		return Tender{}, err
	}

	return tender, nil
}

// ValidationTenderServiceType checks if the Tender type corresponds to the specified values
func ValidationTenderServiceType(serviceType string) error {
	if serviceType == TenderServiceTypeConstruction {
		return nil
	} else if serviceType == TenderServiceTypeDelivery {
		return nil
	} else if serviceType == TenderServiceTypeManufacture {
		return nil
	} else {
		return fmt.Errorf("incorrect service type")
	}
}

// ValidationTenderStatus checks if the tender status corresponds to the specified values
func ValidationTenderStatus(tenderStatus string) error {
	if tenderStatus == TenderStatusCreated {
		return nil
	} else if tenderStatus == TenderStatusPublished {
		return nil
	} else if tenderStatus == TenderStatusClosed {
		return nil
	} else {
		return fmt.Errorf("incorrect status")
	}
}
