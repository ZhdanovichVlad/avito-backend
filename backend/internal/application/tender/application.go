package tender

import (
	"avitoTest/backend/internal/domain/shared"
	"avitoTest/backend/internal/domain/tender"
)

const ServiceTypeIsEmpty = 0
const ServiceTypeNotEmpty = 1
const UsernameSearch = 2

type Application struct {
	TenderInfrastructure tender.TenderInfrastructure
	DataValidator        shared.DataValidator
}
