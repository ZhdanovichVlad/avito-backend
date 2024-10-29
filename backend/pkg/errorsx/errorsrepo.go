package errorsx

import "errors"

var (
	ErrPreparingStatement   = errors.New("error preparing statement")
	ErrExecutingQuery       = errors.New("error executing query")
	ErrInvalidSearchingType = errors.New("unknown searching Type")
	ErrScanRows             = errors.New("database rows scanning error")
)
