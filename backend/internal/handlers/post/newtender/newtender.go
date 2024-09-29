package newtender

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"avitoTest/backend/internal/handlers"
	"avitoTest/backend/internal/lib/api/response"
	"avitoTest/backend/internal/lib/api/typecheck"
	"avitoTest/backend/internal/lib/models"

	"github.com/go-chi/render"
)

// Interface for working with the database
type createTender interface {
	CreateTender(tender models.Tender) (string, error)
	handlers.DataServerChecks
}

// NewTenderH Handler for entering a new tender into the database
func NewTenderH(server createTender) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		const op = "handlers.post.newtender"
		var tender models.Tender

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
		err = tender.Validate()
		if err != nil {
			msgErr := fmt.Errorf("data validation failed. : %w", err)
			response.AnswerError(writer, request, op, http.StatusInternalServerError, msgErr)
			return
		}

		if typecheck.IsTenderServiceTypeIncorrect(tender.ServiceType) {
			msgErr := fmt.Errorf("incorrectly specified tender type")
			response.AnswerError(writer, request, op, http.StatusBadRequest, msgErr)
			return
		}

		сhecking, err := server.CheckUserExists(tender.CreatorUsername)
		if err != nil {
			msgErr := fmt.Errorf("cannot Check User Exists :%s", err)
			response.AnswerError(writer, request, op, http.StatusBadRequest, msgErr)
			return
		}

		if !сhecking {
			msgErr := fmt.Errorf("The user does not exist or is incorrect.")
			response.AnswerError(writer, request, op, http.StatusUnauthorized, msgErr)
			return
		}

		сhecking, err = server.CheckOrganizationExists(tender.OrganizationId)
		if err != nil {
			msgErr := fmt.Errorf("Cannot check organization exists :%s", err)
			response.AnswerError(writer, request, op, http.StatusBadRequest, msgErr)
			return
		}

		if !сhecking {
			msgErr := fmt.Errorf("The organization does not exist or is incorrect.")
			response.AnswerError(writer, request, op, http.StatusUnauthorized, msgErr)
			return
		}

		сhecking, err = server.CheckOrganizationIdAndUserIDExists(tender.OrganizationId, tender.CreatorUsername)
		if err != nil {
			msgErr := fmt.Errorf("cannot Check and Organization User Exists %w", err)
			response.AnswerError(writer, request, op, http.StatusInternalServerError, msgErr)
			return
		}
		if !сhecking {
			msgErr := fmt.Errorf("the link between the employee and the organization could not be verified")
			response.AnswerError(writer, request, op, http.StatusForbidden, msgErr)
			return
		}

		tender.Status = "Created"
		tender.CreatedAt = time.Now()
		tender.Version = 1
		id, err := server.CreateTender(tender)
		if err != nil {
			msgErr := fmt.Errorf("failed to create a new tender %w", err)
			response.AnswerError(writer, request, op, http.StatusInternalServerError, msgErr)
			return
		}

		tender.Id = id
		// The two lines below are needed to make the json format correspond to the output specified in the terms of reference.
		tender.OrganizationId = ""
		tender.CreatorUsername = ""

		render.Status(request, http.StatusOK)
		render.JSON(writer, request, tender)

	}
}
