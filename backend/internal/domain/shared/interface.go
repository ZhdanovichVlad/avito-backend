package shared

// Domain struct with common business logic
type Domain struct {
	infrastructureDB ExistsInfrastructure
}

// Есть валидатор каких-то данных и он умеет делать "полный существует" (UserAndOrgExists) и возвращает статус код http и ошибку

// DataValidator Interface containing common business logic
type DataValidator interface {
	UserAndOrgExists(creatorUsername, organizationId string) (httpCode int, err error)
	LimitAndOffsetCheck(limit, offset string) (limitInt int, offsetInt int, err error)
	IsTenderServiceTypeIncorrect(serviceType string) bool
	UserNameExist(creatorUsername string) (httpCode int, err error)
}

func NewSharedDomain(infra ExistsInfrastructure) *Domain {
	return &Domain{infrastructureDB: infra}
}

// Presentation -> Infrastructure, Application: http -> httpCode
// Infrastructure -> Application, Domain
// Application -> Domain
// Domain -> None
