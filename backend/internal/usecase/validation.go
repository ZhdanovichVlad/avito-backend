package usecase

import (
	"avitoTest/backend/pkg/errorsx"
	"strconv"
)

func LimitAndOffsetValidation(limit, offset string) (limitInt int, offsetInt int, err error) {

	if limit != "" {
		limitInt, err = strconv.Atoi(limit)
		if err != nil {
			return 0, 0, errorsx.ErrInvalidData
		}
	} else {
		limitInt = 5
	}

	if offset != "" {
		offsetInt, err = strconv.Atoi(offset)
		if err != nil {
			return 0, 0, errorsx.ErrInvalidData
		}
	} else {
		offsetInt = 0
	}

	if offsetInt < 0 || limitInt < 0 {
		return 0, 0, errorsx.ErrInvalidData
	}

	return limitInt, offsetInt, nil
}
