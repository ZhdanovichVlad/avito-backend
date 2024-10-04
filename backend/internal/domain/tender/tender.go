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

func NewTender(
	id string,
	name string,
	description string,
	serviceType string,
	status string,
	organizationId string,
	creatorUsername string,
	version int,
	createdAt time.Time,
) (*Tender, error) {
	if len([]rune(t.Name)) > 100 {
		return fmt.Errorf("name len more than 100")
	}
	if len([]rune(t.Description)) > 500 {
		return fmt.Errorf("description len more than 500")
	}
	if len([]rune(t.CreatorUsername)) > 100 {
		return fmt.Errorf("description len more than 500")
	}
	err := t.IsTenderStatusIncorrect()

	return &Tender{}, nil
}

var TenderServiceTypeEnum = [3]string{"Construction", "Delivery", "Manufacture"}
var TenderStatusEnum = [3]string{"Created", "Published", "Closed"}

func (t *Tender) FullValidate() error {
	if len([]rune(t.Name)) > 100 {
		return fmt.Errorf("name len more than 100")
	}
	if len([]rune(t.Description)) > 500 {
		return fmt.Errorf("description len more than 500")
	}
	if len([]rune(t.CreatorUsername)) > 100 {
		return fmt.Errorf("description len more than 500")
	}
	err := t.IsTenderStatusIncorrect()
	if err != nil {
		return err
	}
	err = t.IsTenderServiceTypeIncorrect()
	if err != nil {
		return err
	}

	return nil

}

// IsTenderServiceTypeIncorrect checks if the Tender type corresponds to the specified values
func (t *Tender) IsTenderServiceTypeIncorrect() error {
	for _, enm := range TenderServiceTypeEnum {
		if enm == t.ServiceType {
			return nil
		}
	}
	return fmt.Errorf("incorrect service type")
}

// IsTenderStatusIncorrect checks if the tender status corresponds to the specified values
func (t *Tender) IsTenderStatusIncorrect() error {
	for _, enm := range TenderStatusEnum {
		if enm == t.Status {
			return nil
		}
	}
	return fmt.Errorf("incorrect status")
}
