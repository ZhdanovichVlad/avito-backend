package typecheck

import (
	"avitoTest/backend/internal2/lib/models"
)

// IsTenderServiceTypeIncorrect checks if the Tender type corresponds to the specified values
func IsTenderServiceTypeIncorrect(serviceType string) bool {
	for _, enm := range models.TenderServiceTypeEnum {
		if enm == serviceType {
			return false
		}
	}
	return true
}

// IsTenderStatusIncorrect checks if the tender status corresponds to the specified values
func IsTenderStatusIncorrect(status string) bool {
	for _, enm := range models.TenderStatusEnum {
		if enm == status {
			return false
		}
	}
	return true
}

// IsAuthorTypeEnumIncorrect function for checking the correctness of the status of a user who submits a proposal
func IsAuthorTypeEnumIncorrect(status string) bool {
	for _, enm := range models.AuthorTypeEnum {
		if enm == status {
			return false
		}
	}
	return true
}

// IsAuthorTypeEnumIncorrect check whether the bid status has been correctly transmitted
func IsBidsStatusEmumIncorrect(status string) bool {
	for _, enm := range models.BidsStatusEmum {
		if enm == status {
			return false
		}
	}
	return true
}
