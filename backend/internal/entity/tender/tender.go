package tender

import (
	"fmt"
	"time"
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

var TenderServiceTypeEnum = [3]string{"Construction", "Delivery", "Manufacture"}
var TenderStatusEnum = [3]string{"Created", "Published", "Closed"}

func ValidateTender(t *Tender) error {
	if len([]rune(t.Name)) > 100 && len([]rune(t.Name)) != 0 {
		return fmt.Errorf("name len more than 100 or emty")
	}
	if len([]rune(t.Description)) > 500 && len([]rune(t.Description)) != 0 {
		return fmt.Errorf("description len more than 500 or emty")
	}
	if len([]rune(t.CreatorUsername)) > 100 && len([]rune(t.CreatorUsername)) != 0 {
		return fmt.Errorf("description len more than 500 or emty")
	}
	err := ValidationTenderServiceType(t.ServiceType)
	if err != nil {
		return err
	}

	return nil
}

// ValidationTenderServiceType checks if the Tender type corresponds to the specified values
func ValidationTenderServiceType(serviceType string) error {
	for _, enm := range TenderServiceTypeEnum {
		if enm == serviceType {
			return nil
		}
	}
	return fmt.Errorf("incorrect service type")
}

// ValidationTenderStatus checks if the tender status corresponds to the specified values
func ValidationTenderStatus(tenderStatus string) error {
	for _, enm := range TenderStatusEnum {
		if enm == tenderStatus {
			return nil
		}
	}
	return fmt.Errorf("incorrect status")
}
