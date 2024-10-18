package tender

import (
	"context"

	TenderEntity "avitoTest/backend/internal/entity/tender"
)

type TenderHandlers struct {
	tenderApi TenderHandelrInterface
}

type TenderHandelrInterface interface {
	Create(contex context.Context, input *TenderEntity.Tender) error
}

// NewTenderHandlers Create TenderHandlers
func NewTenderHandlers(handelrInterface TenderHandelrInterface) *TenderHandlers {
	return &TenderHandlers{tenderApi: handelrInterface}
}
