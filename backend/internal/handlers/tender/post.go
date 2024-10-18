package tender

import (
	"avitoTest/backend/pkg/errorsx"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"

	TenderEntity "avitoTest/backend/internal/entity/tender"
)

func (s *TenderHandlers) CreateNewTenderGin(c *gin.Context) {
	const op = "backend.presentation.http.handlers.tenders.post.CreateNewTender"
	context := c.Request.Context()
	var tender TenderEntity.Tender
	err := c.ShouldBindJSON(&tender)
	if err != nil {
		if err.Error() == "EOF" {
			// Обработка ошибки при пустом теле запроса
			c.JSON(http.StatusBadRequest, gin.H{"error": "Request body is empty"})
			return
		}
		// Ответ с ошибкой десериализации
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to deserialize the request: %v", err)})
		return
	}

	err = s.tenderApi.Create(context, &tender)
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

	// Успешный ответ
	c.JSON(http.StatusCreated, tender)
}
