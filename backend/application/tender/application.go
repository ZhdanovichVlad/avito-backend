package tender

import (
	"avitoTest/backend/domain/shared"
	"avitoTest/backend/domain/tender"
)

const ServiceTypeIsEmpty = 0
const ServiceTypeNotEmpty = 1
const UsernameSearch = 2

type NewTenderApplication struct {
	TenderInfrastructure tender.TenderInfrastructure
	DataValidator        shared.DataValidator
}
