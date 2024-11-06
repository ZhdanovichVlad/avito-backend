package application

import (
	"avitoTest/backend/internal/handlers"
	"avitoTest/backend/internal/middleware"
	"avitoTest/backend/pkg/http/ginrouter"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	"net/http"
)

const (
	tenderBasePath      = "/api/tenders"
	PostNewTenderPath   = "/new"
	GetTenderStatusPath = "/:id/status"
	GetMyTendersPath    = "/my"
	PutTenderStatusPath = "/:id/status"
	EditTenderPath      = "/:id/edit"
	PatchTender         = "/:id/rollback/:version"
)

type application struct {
	router *ginrouter.Router
}

func New(router *ginrouter.Router, logger *zap.Logger) *application {
	app := application{router: router}
	app.router.R.Use(middleware.LoggingMiddleware(logger))
	return &app
}

func (a application) RegisterTenderHandlers(h *handlers.TenderHTTP) {

	api := a.router.R.Group(tenderBasePath)
	api.GET("", h.GetAll)
	api.POST(PostNewTenderPath, h.NewTender)
	api.GET(GetTenderStatusPath, h.GetMy)
	api.GET(GetMyTendersPath, h.GetStatus)
	api.PUT(PutTenderStatusPath, h.ChangeStatus)
	api.PATCH(EditTenderPath, h.ChangeTender)
	api.PUT(PatchTender, h.Rollback)
}

func (a application) Run(host string) {
	a.router.R.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ping",
		})
	})
	a.router.R.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Not Found",
		})
	})
	log.Println("server started on", host)
	a.router.R.Run(host)
	log.Println("stopping server")
}
