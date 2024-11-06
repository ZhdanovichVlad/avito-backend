package entityjson

import (
	"avitoTest/backend/pkg/errorsx"
)

type Tender struct {
	Name            string `json:"name,omitempty"`
	Description     string `json:"description,omitempty"`
	ServiceType     string `json:"serviceType,omitempty"`
	OrganizationId  string `json:"organizationId,omitempty"`
	CreatorUsername string `json:"creatorUsername,omitempty"`
}

func (t *Tender) Validate() error {
	const op = "internal.entityjson.Validate"
	if len([]rune(t.Name)) > 100 && len([]rune(t.Name)) != 0 {
		return errorsx.New(errorsx.ErrBadRequest, "name len more than 100 or emty", op, nil)
	}
	if len([]rune(t.Description)) > 500 && len([]rune(t.Description)) != 0 {
		return errorsx.New(errorsx.ErrBadRequest, "description len more than 500 or empty", op, nil)
	}
	if len([]rune(t.CreatorUsername)) > 100 && len([]rune(t.CreatorUsername)) != 0 {
		return errorsx.New(errorsx.ErrBadRequest, "creatorUsername len more than 100 or empty", op, nil)
	}
	return nil
}
