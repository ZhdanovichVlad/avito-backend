package handlers

import (
	"avitoTest/backend/dto"
	"avitoTest/backend/pkg/errorsx"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TenderPresentation struct {
	tenderApi TenderAPI
}

// NewTender create TenderHandlers
func NewTenderPresentation(tenderAPI TenderAPI) *TenderPresentation {
	return &TenderPresentation{tenderApi: tenderAPI}
}

func (s *TenderPresentation) CreateNewGin(c *gin.Context) {

	context := c.Request.Context()
	var tenderDTO dto.TenderHandler
	err := c.ShouldBindJSON(&tenderDTO)
	if err != nil {
		if err.Error() == "EOF" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "request body is empty"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to deserialize the request: %v", err)})
		return
	}

	tenderUseCase, err := s.tenderApi.Create(context, tenderDTO)
	if err != nil {
		switch {
		case errors.Is(err, errorsx.ErrInvalidData):
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		case errors.Is(err, errorsx.ErrInternalRepository):
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal repository error"})
		case errors.Is(err, errorsx.ErrEmployeeNotResponsible):
			c.JSON(http.StatusForbidden, gin.H{"error": "Employee not responsible"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Unable to create a new tender: %v", err)})
		}
		return
	}

	tenderDTO = dto.TenderUseCaseToTenderDTO(tenderUseCase)
	c.JSON(http.StatusCreated, tenderDTO)
}

func (s TenderPresentation) GetAllGin(c *gin.Context) {
	contex := c.Request.Context()
	limit := c.Request.URL.Query().Get("limit")
	offset := c.Request.URL.Query().Get("offset")
	serviceType := c.Request.URL.Query().Get("service_type")

	tenderAPI, err := s.tenderApi.Get(contex, limit, offset, serviceType)
	if err != nil {
		// логика ошибок будет переписана.
		switch {
		case errors.Is(err, errorsx.ErrInvalidData):
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		}
		return
	}

	tenderDTO := make([]dto.TenderHandler, len(tenderAPI))

	for i := 0; i < len(tenderAPI); i++ {
		tenderDTO[i] = dto.TenderUseCaseToTenderDTO(tenderAPI[i])
	}
	c.JSON(http.StatusOK, tenderDTO)
}
