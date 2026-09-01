package handlers

import (
	"avitoTest/backend/internal/entity"
	"avitoTest/backend/internal/entityjson"
	"context"
)

type TenderController interface {
	Create(context context.Context, input entityjson.Tender) (entity.Tender, error)
	Get(context context.Context, limit, offset int, serviceType string) ([]entity.Tender, error)
	GetMy(context context.Context, limit, offset int, userName string) ([]entity.Tender, error)
	GetStatus(context context.Context, tenderId, userName string) (string, error)
	ChangeStatus(context context.Context, tenderId, userName, status string) (entity.Tender, error)
	ChangeTender(context context.Context, tenderId, userName string, input entityjson.Tender) (entity.Tender, error)

	RollbackTender(context context.Context, tenderId, userName string, version int) (entity.Tender, error)
}
