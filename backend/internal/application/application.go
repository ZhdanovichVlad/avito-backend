package application

import (
	TenderHandlers "avitoTest/backend/internal/handlers/tender"
	"avitoTest/backend/pkg/http/ginrouter"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

const (
	tenderBasePath      = "/api/tenders"
	newTenderPath       = "/new"
	tenderStatusPath    = "/{tenderId}/status"
	myTendersPath       = "/my"
	PutTenderStatusPath = "/{tenderId}/status"
)

type application struct {
	router *ginrouter.Router
}

func New(router *ginrouter.Router) *application {
	app := application{router: router}
	return &app
}

func (a application) RegisterTenderHandlers(h *TenderHandlers.TenderHandlers) {

	api := a.router.R.Group(tenderBasePath)
	{
		api.POST(newTenderPath, h.CreateNewTenderGin)
	}
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
