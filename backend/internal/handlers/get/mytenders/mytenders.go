package mytenders

import (
	"errors"
	"fmt"
	"net/http"

	"avitoTest/backend/internal/handlers/get"
	"avitoTest/backend/internal/lib/api/limitandoffsetcheck"
	"avitoTest/backend/internal/lib/api/response"
	"avitoTest/backend/internal/lib/models"

	"github.com/go-chi/render"
)

// GetTasksH Handler to retrieve a list of tenders according to specified conditions. The нandler supports pagination
func GetMyTender(server get.ServerGet) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		const op = "backend.internal.handlers.get.getMyTender"

		reqQuery := request.URL.Query()
		limit := reqQuery.Get("limit")
		offset := reqQuery.Get("offset")
		username := reqQuery.Get("username")

		limitInt, offsetInt, err := limitandoffsetcheck.LimitAndOffsetCheck(limit, offset)
		if err != nil {
			err = errors.New("connot convert limit or offset to int")
			response.AnswerError(writer, request, op, http.StatusBadRequest, err)
			return
		}

		сhecking, err := server.CheckUserExists(username)
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

		tenders := make([]models.Tender, 0, limitInt)

		err = server.GetTenders(&tenders, limitInt, offsetInt, username, get.UsernameSearch)
		if err != nil {
			err = fmt.Errorf("error when receiving data from the server. %s", err)
			response.AnswerError(writer, request, op, http.StatusInternalServerError, err)
			return
		}

		render.Status(request, http.StatusOK)
		render.JSON(writer, request, tenders)
	}
}
