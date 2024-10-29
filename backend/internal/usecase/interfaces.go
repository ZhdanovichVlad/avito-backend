package usecase

import (
	"avitoTest/backend/internal/entity"
	"context"
)

type TenderUseCaseInterface interface {
	CreateTender(context context.Context, tender *entity.Tender) (string, error)
	GetTenders(context context.Context, limit, offset int, searchInfo string, searchingType int) ([]entity.Tender, error)
	ValidateResponsibleEmployee(context context.Context, organizationId, creatorUsername string) (bool, error)
	//GetFullTender(tender *tender.Tender, tenderId string) error
	//UpdateTenderStatus(tenderId string, status string) (err error)
}
