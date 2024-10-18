package errorsx

import "errors"

var (
	ErrInvalidData            = errors.New("invalid data")
	ErrEmployeeNotResponsible = errors.New("the employee is not responsible for this organization")
	ErrInternalRepository     = errors.New("error while retrieving information from the database")
)

type ErrorAnswer struct {
	ErrorMsg string `json:"Error,omitempty"`
}
