package usecase

import (
	"avitoTest/backend/internal/entity"
	"avitoTest/backend/internal/handlers"
	"avitoTest/backend/pkg/errorsx"
	"context"
)

const (
	ServiceTypeIsEmpty  = 0
	ServiceTypeNotEmpty = 1
	UsernameSearch      = 2
)

type TenderApplacation struct {
	tenderRepository TenderUseCaseInterface
	//DataValidator        shared.DataValidator
}

func New(repo TenderUseCaseInterface) *TenderApplacation {
	return &TenderApplacation{tenderRepository: repo}
}

// Create method for creating a new tender
func (server *TenderApplacation) Create(context context.Context, input handlers.TenderDTO) (entity.Tender, error) {
	tender, err := entity.NewTenderUseCase(input)
	if err != nil {
		return entity.Tender{}, errorsx.ErrInvalidData
	}

	exist, err := server.tenderRepository.ValidateResponsibleEmployee(context, tender.OrganizationId, tender.CreatorUsername)
	if err != nil {
		return entity.Tender{}, errorsx.ErrInternalRepository
	}
	if !exist {
		return entity.Tender{}, errorsx.ErrEmployeeNotResponsible
	}

	id, err := server.tenderRepository.CreateTender(context, &tender)
	if err != nil {
		return entity.Tender{}, errorsx.ErrInternalRepository
	}

	tender.Id = id
	return tender, nil
}

func (server *TenderApplacation) Get(context context.Context, limit, offset, serviceInfo string) ([]entity.Tender, error) {

	limitInt, offsetInt, err := LimitAndOffsetValidation(limit, offset)
	if err != nil {
		return nil, errorsx.ErrInvalidData
	}

	if serviceInfo != "" {
		err = entity.ValidationTenderServiceType(serviceInfo)
		if err != nil {
			return nil, errorsx.ErrInvalidData
		}
	}

	if serviceInfo != "" {
		err = entity.ValidationTenderServiceType(serviceInfo)
		if err != nil {
			return nil, errorsx.ErrInvalidServiceType
		}
	}

	var searchingType int
	if serviceInfo == "" {
		searchingType = ServiceTypeIsEmpty
	} else {
		searchingType = ServiceTypeNotEmpty
	}

	tenders, err := server.tenderRepository.GetTenders(context, limitInt, offsetInt, serviceInfo, searchingType)
	if err != nil {
		return nil, errorsx.ErrInternalRepository
	}
	return tenders, nil
}
