package tender

// Interface containing common tender business logic
type TenderDomain struct {
	infrastructureDB TenderInfrastructure
}

type TenderInfrastructure interface {
	CreateTender(tender *Tender) (string, error)
	GetTenders(tenders *[]Tender, limit, offset int, searchInfo string, serchingType int) error
	GetFullTender(tender *Tender, tenderId string) error
	UpdateTenderStatus(tenderId string, status string) (err error)
}

func NewSharedDonain(infastructure TenderInfrastructure) *TenderDomain {
	return &TenderDomain{infrastructureDB: infastructure}
}
