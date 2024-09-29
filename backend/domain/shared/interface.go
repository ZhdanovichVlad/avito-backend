package shared

// struct with common business logic
type ShardeDomain struct {
	infrastructureDB ExistsInfrastructure
}

// Interface containing common business logic
type DataValidator interface {
	FullExist(creatorUsername, organizationId string) (httpCode int, err error)
	LimitAndOffsetCheck(limit, offset string) (limitInt int, offsetInt int, err error)
	IsTenderServiceTypeIncorrect(serviceType string) bool
	UserNameExist(creatorUsername string) (httpCode int, err error)
}

func NewSharedDonain(infastructure ExistsInfrastructure) *ShardeDomain {
	return &ShardeDomain{infrastructureDB: infastructure}
}
