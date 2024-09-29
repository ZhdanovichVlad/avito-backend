package tenders

import (
	tenderApplication "avitoTest/backend/application/tender"
	tenderDomain "avitoTest/backend/domain/tender"
	"avitoTest/backend/internal/lib/api/response"
	"errors"
	"fmt"
	"github.com/go-chi/render"
	"io"
	"net/http"
)

func CreateNewTender(s tenderApplication.NewTenderApplication) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		const op = "backend.interfaces.http.handlers.tenders.postnewtedner"
		var tender tenderDomain.Tender

		err := render.DecodeJSON(request.Body, &tender)
		if errors.Is(err, io.EOF) {
			msgErr := fmt.Errorf("handle an error if we receive a request with an empty body")
			response.AnswerError(writer, request, op, http.StatusBadRequest, msgErr)
			return
		}
		if err != nil {
			msgErr := fmt.Errorf("failed to deserialize the request. : %w", err)
			response.AnswerError(writer, request, op, http.StatusInternalServerError, msgErr)
			return
		}

		httpCode, err := s.NewTender(&tender)
		if err != nil {
			msgErr := fmt.Errorf("Unable to create a new tender: %w", err)
			response.AnswerError(writer, request, op, httpCode, msgErr)
			return
		}

		render.Status(request, http.StatusCreated)
		render.JSON(writer, request, tender)
	}
}
