package changebidstatus

import (
	"avitoTest/backend/internal2/lib/api/typecheck"
	"fmt"
	"net/http"

	"avitoTest/backend/internal2/handlers/get"
	"avitoTest/backend/internal2/lib/api/response"
	"avitoTest/backend/internal2/lib/models"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type bidStatusChangeI interface {
	GetBid(bid *models.Bid, bidID string) (successfulRequest bool, err error)
	GetCompanyIDbyUser(username string) (companyId string, err error)
	CheckBidExists(bidID string) (bool, error)
	UpdateBidStatus(bidID, status string) error
	get.ServerGet
}

// BidStatusChange bid status change method
func BidStatusChange(server bidStatusChangeI) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		const op = "backend.internal2.handlers.put.BidStatusChange"
		bidId := chi.URLParam(request, "bidId")
		username := request.URL.Query().Get("username")
		status := request.URL.Query().Get("status")

		if typecheck.IsBidsStatusEmumIncorrect(status) {
			msgErr := fmt.Errorf("incorrect bid status")
			response.AnswerError(writer, request, op, http.StatusBadRequest, msgErr)
			return
		}

		successfulRequest, err := server.CheckBidExists(bidId)
		if err != nil {
			msgErr := fmt.Errorf("failed to retrieve information from the database %w", err)
			response.AnswerError(writer, request, op, http.StatusInternalServerError, msgErr)
			return
		}
		if !successfulRequest {
			msgErr := fmt.Errorf("couldn't find the bid")
			response.AnswerError(writer, request, op, http.StatusNotFound, msgErr)
			return
		}

		bid := models.Bid{}
		successfulRequest, err = server.GetBid(&bid, bidId)
		if err != nil {
			msgErr := fmt.Errorf("failed to retrieve information from the database %w", err)
			response.AnswerError(writer, request, op, http.StatusInternalServerError, msgErr)
			return
		}
		if !successfulRequest {
			msgErr := fmt.Errorf("couldn't find the bid")
			response.AnswerError(writer, request, op, http.StatusNotFound, msgErr)
			return
		}

		var organizationId string
		switch bid.AuthorType {
		case models.AuthorTypeEnum[0]: // user
			cheking, err := server.CheckUserExists(username)
			if err != nil {
				msgErr := fmt.Errorf("cannot check user exists %w", err)
				response.AnswerError(writer, request, op, http.StatusInternalServerError, msgErr)
				return
			}
			if !cheking {
				msgErr := fmt.Errorf("the user does not exist or is incorrect.")
				response.AnswerError(writer, request, op, http.StatusUnauthorized, msgErr)
				return
			}
			organizationId, err = server.GetCompanyIDbyUser(username)
			if err != nil {
				msgErr := fmt.Errorf("failed to retrieve organization data :%s", err)
				response.AnswerError(writer, request, op, http.StatusInternalServerError, msgErr)
				return
			}
			if bid.AuthorId != organizationId && bid.AuthorId != username {
				msgErr := fmt.Errorf("this bid is not available to you")
				response.AnswerError(writer, request, op, http.StatusForbidden, msgErr)
				return
			}
		case models.AuthorTypeEnum[1]: // organization
			сhecking, err := server.CheckOrganizationExists(username)
			if err != nil {
				msgErr := fmt.Errorf("cannot check organization exists %w", err)
				response.AnswerError(writer, request, op, http.StatusInternalServerError, msgErr)
				return
			}
			if !сhecking {
				msgErr := fmt.Errorf("this user cannot get information for this tender")
				response.AnswerError(writer, request, op, http.StatusForbidden, msgErr)
				return
			}
		}

		err = server.UpdateBidStatus(bidId, status)
		if err != nil {
			msgErr := fmt.Errorf("cannot update bid status %w", err)
			response.AnswerError(writer, request, op, http.StatusInternalServerError, msgErr)
			return
		}

		bid.Status = status

		render.Status(request, http.StatusOK)
		render.JSON(writer, request, bid)
	}
}
