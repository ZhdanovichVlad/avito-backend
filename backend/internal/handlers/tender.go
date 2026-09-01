package handlers

import (
	"avitoTest/backend/internal/entityjson"
	"avitoTest/backend/pkg/errorsx"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TenderHTTP struct {
	tenderController TenderController
}

// NewTenderController
func NewTenderController(controller TenderController) *TenderHTTP {
	return &TenderHTTP{tenderController: controller}
}

func (t *TenderHTTP) NewTender(c *gin.Context) {
	const op = "internal.handlers.NewTender"
	context := c.Request.Context()
	var tenderJSON entityjson.Tender
	err := c.ShouldBindJSON(&tenderJSON)
	if err != nil {
		if err.Error() == "EOF" {
			errx := errorsx.New(errorsx.ErrBadRequest, "request body is empty", op, err)
			c.JSON(errx.ErrCodeToHTTPStatus(), errx.Message)
			return
		}
		errx := errorsx.New(errorsx.ErrInternal, "failed to deserialize the request", op, err)
		c.JSON(errx.ErrCodeToHTTPStatus(), errx.Message)
		return
	}

	tenderEntity, err := t.tenderController.Create(context, tenderJSON)
	if err != nil {
		var errx *errorsx.Errorx
		if errors.As(err, &errx) {
			c.JSON(errx.ErrCodeToHTTPStatus(), errx.Message)
			return
		} else {
			errx := errorsx.New(errorsx.ErrInternal, "unable to create a new tender", op, err)
			c.JSON(errx.ErrCodeToHTTPStatus(), errx.Message)
			return
		}

	}
	c.JSON(http.StatusCreated, tenderEntity)
}

func (t *TenderHTTP) GetAll(c *gin.Context) {
	const op = "internal.handlers.GetAll"
	contex := c.Request.Context()
	limit := c.Request.URL.Query().Get("limit")
	offset := c.Request.URL.Query().Get("offset")
	serviceType := c.Request.URL.Query().Get("service_type")

	limitInt, offsetInt, err := LimitAndOffsetValidation(limit, offset)
	if err != nil {
		var errx *errorsx.Errorx
		if errors.As(err, &errx) {
			c.JSON(errx.ErrCodeToHTTPStatus(), errx.Message)
			return
		} else {
			errx := errorsx.New(errorsx.ErrInternal, "unable to get tenders", op, err)
			c.JSON(errx.ErrCodeToHTTPStatus(), errx.Message)
			return
		}
	}

	tenders, err := t.tenderController.Get(contex, limitInt, offsetInt, serviceType)
	if err != nil {
		var errx *errorsx.Errorx
		if errors.As(err, &errx) {
			c.JSON(errx.ErrCodeToHTTPStatus(), errx.Message)
			return
		} else {
			errx := errorsx.New(errorsx.ErrInternal, "unable to create a new tender", op, err)
			c.JSON(errx.ErrCodeToHTTPStatus(), errx.Message)
			return
		}
	}

	c.JSON(http.StatusOK, tenders)
}

func (t *TenderHTTP) GetMy(c *gin.Context) {
	const op = "internal.handlers.Get"
	contex := c.Request.Context()
	limit := c.Request.URL.Query().Get("limit")
	offset := c.Request.URL.Query().Get("offset")
	username := c.Request.URL.Query().Get("username")

	limitInt, offsetInt, err := LimitAndOffsetValidation(limit, offset)
	if err != nil {
		var errx *errorsx.Errorx
		if errors.As(err, &errx) {
			c.JSON(errx.ErrCodeToHTTPStatus(), errx.Message)
			return
		} else {
			errx := errorsx.New(errorsx.ErrInternal, "unable to get limit and offset", op, err)
			c.JSON(errx.ErrCodeToHTTPStatus(), errx.Message)
			return
		}
	}

	tenders, err := t.tenderController.GetMy(contex, limitInt, offsetInt, username)
	if err != nil {
		var errx *errorsx.Errorx
		if errors.As(err, &errx) {
			c.JSON(errx.ErrCodeToHTTPStatus(), errx.Message)
			return
		} else {
			errx := errorsx.New(errorsx.ErrInternal, "unable to create a Get tenders", op, err)
			c.JSON(errx.ErrCodeToHTTPStatus(), errx.Message)
			return
		}

	}
	c.JSON(http.StatusOK, tenders)
}

func (t *TenderHTTP) GetStatus(c *gin.Context) {
	const op = "handler.tender.GetStatus"
	context := c.Request.Context()
	tenderID := c.Param("id")
	userName := c.Request.URL.Query().Get("username")

	status, err := t.tenderController.GetStatus(context, tenderID, userName)
	if err != nil {
		var errx *errorsx.Errorx
		if errors.As(err, &errx) {
			c.JSON(errx.ErrCodeToHTTPStatus(), errx.Message)
			return
		} else {
			errx := errorsx.New(errorsx.ErrInternal, "unable to get status", op, err)
			c.JSON(errx.ErrCodeToHTTPStatus(), errx.Message)
			return
		}
	}

	c.JSON(http.StatusOK, status)
}

func (t *TenderHTTP) ChangeStatus(c *gin.Context) {
	const op = "handler.tender.ChangeStatus"
	context := c.Request.Context()
	tenderID := c.Param("id")
	userName := c.Request.URL.Query().Get("username")
	status := c.Request.URL.Query().Get("status")

	tender, err := t.tenderController.ChangeStatus(context, tenderID, userName, status)
	if err != nil {
		var errx *errorsx.Errorx
		if errors.As(err, &errx) {
			c.JSON(errx.ErrCodeToHTTPStatus(), errx.Message)
			return
		} else {
			errx := errorsx.New(errorsx.ErrInternal, "unable to change status", op, err)
			c.JSON(errx.ErrCodeToHTTPStatus(), errx.Message)
			return
		}
	}
	c.JSON(http.StatusOK, tender)
}

func (t *TenderHTTP) ChangeTender(c *gin.Context) {
	const op = "handler.tender.ChangeTender"
	context := c.Request.Context()
	tenderID := c.Param("id")
	userName := c.Request.URL.Query().Get("username")
	var tenderJSON entityjson.Tender

	err := c.ShouldBindJSON(&tenderJSON)
	if err != nil {
		if err.Error() == "EOF" {
			errx := errorsx.New(errorsx.ErrBadRequest, "request body is empty", op, err)
			c.JSON(errx.ErrCodeToHTTPStatus(), errx.Message)
			return
		}
		errx := errorsx.New(errorsx.ErrInternal, "failed to deserialize the request", op, err)
		c.JSON(errx.ErrCodeToHTTPStatus(), errx.Message)
		return
	}

	if len([]rune(tenderJSON.Description)) > 500 && len([]rune(tenderJSON.Description)) != 0 {
		errx := errorsx.New(errorsx.ErrInternal, "description len more than 500 or empty", op, nil)
		c.JSON(errx.ErrCodeToHTTPStatus(), errx.Message)
	}

	tender, err := t.tenderController.ChangeTender(context, tenderID, userName, tenderJSON)
	if err != nil {
		var errx *errorsx.Errorx
		if errors.As(err, &errx) {
			c.JSON(errx.ErrCodeToHTTPStatus(), errx.Message)
			return
		} else {
			errx := errorsx.New(errorsx.ErrInternal, "unable to change status", op, err)
			c.JSON(errx.ErrCodeToHTTPStatus(), errx.Message)
			return
		}
	}
	c.JSON(http.StatusOK, tender)
}

func (t *TenderHTTP) Rollback(g *gin.Context) {
	const op = "handler.tender.RollbackTender"

	context := g.Request.Context()
	tenderID := g.Param("id")
	tenderVersion := g.Param("version")
	userName := g.Request.URL.Query().Get("username")

	tenderVersionInt, err := ValidateAndConvertVersion(tenderVersion)
	if err != nil {
		var errx *errorsx.Errorx
		if errors.As(err, &errx) {
			g.JSON(errx.ErrCodeToHTTPStatus(), errx.Message)
			return
		} else {
			errx := errorsx.New(errorsx.ErrInternal, "unable to change tender", op, err)
			g.JSON(errx.ErrCodeToHTTPStatus(), errx.Message)
			return
		}
	}

	tender, err := t.tenderController.RollbackTender(context, tenderID, userName, tenderVersionInt)
	if err != nil {
		var errx *errorsx.Errorx
		if errors.As(err, &errx) {
			g.JSON(errx.ErrCodeToHTTPStatus(), errx.Message)
			return
		} else {
			errx := errorsx.New(errorsx.ErrInternal, "unable to change tender", op, err)
			g.JSON(errx.ErrCodeToHTTPStatus(), errx.Message)
			return
		}
	}

	g.JSON(http.StatusOK, tender)

}
