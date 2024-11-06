package errorsx

import (
	"fmt"
	"net/http"
)

// list of errors in the application
const (
	ErrNotFound   = "notfound"
	ErrInternal   = "internal"
	ErrBadRequest = "bad request"
	ErrForbidden  = "the employee is not responsible for this organization"
)

var codeToHTTPStatusMap = map[string]int{
	ErrNotFound:   http.StatusNotFound,
	ErrInternal:   http.StatusInternalServerError,
	ErrBadRequest: http.StatusBadRequest,
	ErrForbidden:  http.StatusForbidden,
	// Другие соответствия кодов ошибок и HTTP-статусов
}

type Errorx struct {
	//user error
	Code string `json:"code"`
	//clear message to the user
	Message string `json:"message"`
	//the location where the error occurred
	Op string `json:"op"`
	//standard error
	Err error `json:"err"`
}

func New(code string, message string, op string, err error) *Errorx {
	return &Errorx{
		Code:    code,
		Message: message,
		Op:      op,
		Err:     err,
	}
}

func (e *Errorx) ErrCodeToHTTPStatus() int {
	if v, ok := codeToHTTPStatusMap[e.Code]; ok {
		return v
	}
	return http.StatusInternalServerError
}

func (e *Errorx) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %s: %v", e.Op, e.Message, e.Err)
	}
	return e.Message
}
