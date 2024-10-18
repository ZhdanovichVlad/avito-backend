package tender

import (
	"avitoTest/backend/internal/entity/tender"
	"context"
)

const ServiceTypeIsEmpty = 0
const ServiceTypeNotEmpty = 1
const UsernameSearch = 2

type TenderApplacation struct {
	tenderRepository TenderUseCaseInterface
	//DataValidator        shared.DataValidator
}

type TenderUseCaseInterface interface {
	CreateTender(context context.Context, tender *tender.Tender) (string, error)
	ValidateResponsibleEmployee(organizationId, creatorUsername string) (bool, error)
	//GetTenders(tenders *[]tender.Tender, limit, offset int, searchInfo string, serchingType int) error
	//GetFullTender(tender *tender.Tender, tenderId string) error
	//UpdateTenderStatus(tenderId string, status string) (err error)
}

func New(repo TenderUseCaseInterface) *TenderApplacation {
	return &TenderApplacation{tenderRepository: repo}
}
