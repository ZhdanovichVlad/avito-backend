package bidstatus

import (
	"fmt"
	"net/http"

	"avitoTest/backend/internal2/handlers/get"
	"avitoTest/backend/internal2/lib/api/response"
	"avitoTest/backend/internal2/lib/models"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type bidStatusI interface {
	GetBid(bid *models.Bid, bidID string) (successfulRequest bool, err error)
	GetCompanyIDbyUser(username string) (companyId string, err error)
	CheckBidExists(bidID string) (bool, error)
	get.ServerGet
}

// BidStatus method for obtaining bid status
func BidStatus(server bidStatusI) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		const op = "backend.internal2.handlers.get.bidStatus"
		bidId := chi.URLParam(request, "bidId")
		username := request.URL.Query().Get("username")

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
		render.Status(request, http.StatusOK)
		render.JSON(writer, request, bid.Status)
	}
}
