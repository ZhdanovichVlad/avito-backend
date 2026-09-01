package usecase

import (
	"avitoTest/backend/internal/entity"
	"context"
)

type TenderRepository interface {
	CreateTender(context context.Context, tender *entity.Tender) (string, error)
	GetTenders(context context.Context, limit, offset int) ([]entity.Tender, error)
	GetUserTenders(context context.Context, limit, offset int, user string) ([]entity.Tender, error)
	GetTendersWithServiceType(context context.Context, limit, offset int, ServiceType string) ([]entity.Tender, error)
	GetTenderStatus(context context.Context, id string) (string, string, error)
	GetTender(context context.Context, tenderId string) (entity.Tender, error)
	UpdateTenderStatus(context context.Context, tenderId string, status string) (string, error)
	UpdateTender(context context.Context, tender *entity.Tender) error
	RollBackTender(context context.Context, tenderID string, version int) error

	ValidateResponsibleEmployee(context context.Context, organizationId, creatorUsername string) error
	CheckUserExists(context context.Context, creatorUsername string) error
}
