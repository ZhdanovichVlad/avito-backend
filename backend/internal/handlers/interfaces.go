package handlers

import (
	"avitoTest/backend/dto"
	"avitoTest/backend/internal/entity"
	"context"
)

type TenderAPI interface {
	Create(context context.Context, input dto.TenderHandler) (entity.Tender, error)
	Get(context context.Context, limit, offset, serviceType string) ([]entity.Tender, error)
}
