package shared

import (
	"fmt"
	"net/http"
	"strconv"
)

// ExistsInfrastructure An interface that verifies the existence of the organization and employees
type ExistsInfrastructure interface {
	OrganizationExists(organizationId string) (bool, error)
	CheckUserExistsByID(userID string) (bool, error)
	CheckUserExists(creatorUsername string) (bool, error)
	CheckOrganizationIdAndUserIDExists(organizationId, creatorUsername string) (bool, error)
}

var TenderServiceTypeEnum = [3]string{"Construction", "Delivery", "Manufacture"}

// UserAndOrgExists function for checking the existence of a user, company and whether the user is a representative of the company
func (s *Domain) UserAndOrgExists(creatorUsername, organizationId string) (httpCode int, err error) {
	httpCode, err = s.UserNameExist(creatorUsername)
	if err != nil {
		return
	}

	сhecking, err := s.infrastructureDB.OrganizationExists(organizationId)
	if err != nil {
		msgErr := fmt.Errorf("Cannot check organization exists :%s", err)
		return http.StatusNotFound, msgErr
	}

	if !сhecking {
		msgErr := fmt.Errorf("The organization does not exist or is incorrect.")
		return http.StatusUnauthorized, msgErr
	}

	сhecking, err = s.infrastructureDB.CheckOrganizationIdAndUserIDExists(organizationId, creatorUsername)
	if err != nil {
		msgErr := fmt.Errorf("cannot Check and Organization User Exists %w", err)
		return http.StatusNotFound, msgErr
	}
	if !сhecking {
		msgErr := fmt.Errorf("the link between the employee and the organization could not be verified")
		return http.StatusUnauthorized, msgErr
	}

	return http.StatusOK, err
}

func (s *Domain) UserNameExist(creatorUsername string) (httpCode int, err error) {
	const op = "application.shared"
	сhecking, err := s.infrastructureDB.CheckUserExists(creatorUsername)
	if err != nil {
		msgErr := fmt.Errorf("op - %s. cannot check user exists :%s", op, err)
		return http.StatusNotFound, msgErr
	}

	if !сhecking {
		msgErr := fmt.Errorf("The user does not exist or is incorrect.")
		return http.StatusUnauthorized, msgErr
	}
	return http.StatusOK, nil
}

// LimitAndOffsetCheck function for checking limit and offset for int data content
func (s *Domain) LimitAndOffsetCheck(limit, offset string) (limitInt int, offsetInt int, err error) {
	const op = "lib.api.LimitAndOffsetCheck"

	if limit != "" {
		limitInt, err = strconv.Atoi(limit)
		if err != nil {
			msgErr := fmt.Errorf("error in  %s. Error %s", op, err)
			return 0, 0, msgErr
		}
	} else {
		limitInt = 5
	}

	if offset != "" {
		offsetInt, err = strconv.Atoi(offset)
		if err != nil {
			msgErr := fmt.Errorf("error in  %s. Error %s", op, err)
			return 0, 0, msgErr
		}
	} else {
		offsetInt = 0
	}

	if offsetInt < 0 || limitInt < 0 {
		msgErr := fmt.Errorf("limit or offset is negative", op)
		return 0, 0, msgErr
	}

	return limitInt, offsetInt, nil
}

func (s *Domain) IsTenderServiceTypeIncorrect(serviceType string) bool {
	for _, enm := range TenderServiceTypeEnum {
		if enm == serviceType {
			return false
		}
	}
	return true
}
