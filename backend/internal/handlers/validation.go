package handlers

import (
	"avitoTest/backend/pkg/errorsx"
	"strconv"
)

func LimitAndOffsetValidation(limit, offset string) (limitInt int, offsetInt int, err error) {
	const op = "handlers.validation.LimitAndOffsetValidation"
	if limit != "" {
		limitInt, err = strconv.Atoi(limit)
		if err != nil {
			return 0, 0, errorsx.New(errorsx.ErrBadRequest, "invalid limit", op, err)
		}
	} else {
		limitInt = 5
	}

	if offset != "" {
		offsetInt, err = strconv.Atoi(offset)
		if err != nil {
			return 0, 0, errorsx.New(errorsx.ErrBadRequest, "invalid offset", op, err)
		}
	} else {
		offsetInt = 0
	}

	if offsetInt < 0 || limitInt < 0 {
		return 0, 0, errorsx.New(errorsx.ErrBadRequest, "offset or limit les then zero", op, nil)
	}

	return limitInt, offsetInt, nil
}

func ValidateAndConvertVersion(versionStr string) (int, error) {
	const op = "handlers.validation.ValidateVersion"

	versionInt, err := strconv.Atoi(versionStr)
	if err != nil {
		return 0, errorsx.New(errorsx.ErrBadRequest, "invalid version", op, nil)
	}
	if versionInt <= 0 {
		return 0, errorsx.New(errorsx.ErrBadRequest, "invalid version", op, nil)
	}
	return versionInt, nil
}
