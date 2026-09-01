package usecase

import (
	"context"

	"avitoTest/backend/internal/entity"
	"avitoTest/backend/internal/entityjson"
)

type TenderService struct {
	tenderRepository TenderRepository
}

func New(repo TenderRepository) *TenderService {
	return &TenderService{tenderRepository: repo}
}

// Create method for creating a new tender
func (t *TenderService) Create(context context.Context, tenderJson entityjson.Tender) (entity.Tender, error) {
	tender := entity.NewTender(tenderJson)

	err := entity.ValidationTenderServiceType(tender.ServiceType)
	if err != nil {
		return entity.Tender{}, err
	}

	err = t.tenderRepository.ValidateResponsibleEmployee(context, tender.OrganizationId, tender.CreatorUsername)
	if err != nil {
		return entity.Tender{}, err
	}

	id, err := t.tenderRepository.CreateTender(context, tender)
	if err != nil {
		return entity.Tender{}, err
	}

	tender.Id = id
	return *tender, nil
}

func (t *TenderService) Get(context context.Context, limit, offset int, serviceInfo string) ([]entity.Tender, error) {
	const op = "usecase.tender.Get"

	var err error
	var tenders []entity.Tender

	switch serviceInfo {
	case "":
		tenders, err = t.tenderRepository.GetTenders(context, limit, offset)
		if err != nil {
			return nil, err
		}
	default:
		err = entity.ValidationTenderServiceType(serviceInfo)
		if err != nil {
			return nil, err
		}

		tenders, err = t.tenderRepository.GetTenders(context, limit, offset)
		if err != nil {
			return nil, err
		}
	}

	return tenders, nil
}

func (t *TenderService) GetMy(context context.Context, limit, offset int, userName string) ([]entity.Tender, error) {
	const op = "usecase.tender.Get"

	var err error
	var tenders []entity.Tender

	err = t.tenderRepository.CheckUserExists(context, userName)
	if err != nil {
		return nil, err
	}

	tenders, err = t.tenderRepository.GetUserTenders(context, limit, offset, userName)
	if err != nil {
		return nil, err
	}

	return tenders, nil
}

func (t *TenderService) GetStatus(context context.Context, tenderId, userName string) (string, error) {
	const op = "usecase.tender.GetStatus"

	err := t.tenderRepository.CheckUserExists(context, userName)
	if err != nil {
		return "", nil
	}

	status, companyid, err := t.tenderRepository.GetTenderStatus(context, tenderId)

	err = t.tenderRepository.ValidateResponsibleEmployee(context, companyid, userName)
	if err != nil {
		return "", nil
	}

	return status, nil
}

func (t *TenderService) ChangeStatus(context context.Context, tenderId, userName, status string) (entity.Tender, error) {
	const op = "usecase.tender.GetStatus"
	err := entity.ValidationTenderStatus(status)
	if err != nil {
		return entity.Tender{}, err
	}

	err = t.tenderRepository.CheckUserExists(context, userName)
	if err != nil {
		return entity.Tender{}, err
	}

	tender, err := t.tenderRepository.GetTender(context, tenderId)
	if err != nil {
		return entity.Tender{}, err
	}

	err = t.tenderRepository.ValidateResponsibleEmployee(context, tender.OrganizationId, userName)
	if err != nil {
		return entity.Tender{}, err
	}

	newStatus, err := t.tenderRepository.UpdateTenderStatus(context, tenderId, status)
	if err != nil {
		return entity.Tender{}, err
	}

	tender.Status = newStatus
	return tender, nil
}

func (t *TenderService) ChangeTender(context context.Context, tenderId, userName string, tenderJson entityjson.Tender) (entity.Tender, error) {
	const op = "usecase.tender.ChangeTender"

	tenderOld, err := t.tenderRepository.GetTender(context, tenderId)
	if err != nil {
		return entity.Tender{}, nil
	}
	err = t.tenderRepository.ValidateResponsibleEmployee(context, tenderOld.OrganizationId, userName)
	if err != nil {
		return entity.Tender{}, nil
	}

	tenderNew := entity.NewTenderForUpdate(tenderJson)

	err = entity.ValidationTenderServiceType(tenderNew.ServiceType)
	if err != nil {
		return entity.Tender{}, err
	}

	tenderNew.Id = tenderId
	err = t.tenderRepository.UpdateTender(context, tenderNew)
	if err != nil {
		return entity.Tender{}, err
	}

	tender, err := t.tenderRepository.GetTender(context, tenderId)
	if err != nil {
		return entity.Tender{}, err
	}

	return tender, nil
}

func (t *TenderService) RollbackTender(context context.Context, tenderId, userName string, version int) (entity.Tender, error) {
	const op = "usecase.tender.ChangeTender"

	tender, err := t.tenderRepository.GetTender(context, tenderId)
	if err != nil {
		return entity.Tender{}, nil
	}
	err = t.tenderRepository.ValidateResponsibleEmployee(context, tender.OrganizationId, userName)
	if err != nil {
		return entity.Tender{}, nil
	}
	err = t.tenderRepository.RollBackTender(context, tenderId, version)
	if err != nil {
		return entity.Tender{}, nil
	}

	tender, err = t.tenderRepository.GetTender(context, tenderId)
	if err != nil {
		return entity.Tender{}, nil
	}

	return tender, nil
}
