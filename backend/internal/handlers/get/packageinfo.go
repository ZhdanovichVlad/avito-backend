package get

import (
	"avitoTest/backend/internal/handlers"
	"avitoTest/backend/internal/lib/models"
)

// description of database search options
const ServiceTypeIsEmpty = 0
const ServiceTypeNotEmpty = 1
const UsernameSearch = 2

const BidSearchByUser = 0
const BidSearchByUserAncCompanyID = 1
const BidSearchTenderId = 2

// description of server methods for working with get handles
type ServerGet interface {
	GetTenders(tenders *[]models.Tender, limit, offset int, searchInfo string, serchingType int) error
	handlers.DataServerChecks
}
