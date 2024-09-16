package router

import (
	"avitoTest/backend/internal/handlers/get/bidstatus"
	"avitoTest/backend/internal/handlers/get/mybids"
	"avitoTest/backend/internal/handlers/get/mytenders"
	"avitoTest/backend/internal/handlers/get/ping"
	"avitoTest/backend/internal/handlers/get/tenderbids"
	"avitoTest/backend/internal/handlers/get/tenders"
	"avitoTest/backend/internal/handlers/get/tenderstatus"
	"avitoTest/backend/internal/handlers/patch/editbid"
	"avitoTest/backend/internal/handlers/patch/edittender"
	"avitoTest/backend/internal/handlers/post/newbid"
	"avitoTest/backend/internal/handlers/post/newtender"
	"avitoTest/backend/internal/handlers/put/changebidstatus"
	"avitoTest/backend/internal/handlers/put/changetenderstatus"
	"avitoTest/backend/internal/handlers/put/rollbackbid"
	"avitoTest/backend/internal/handlers/put/rollbacktender"
	logmiddleware "avitoTest/backend/internal/middleware/log"
	"avitoTest/backend/internal/storage"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// NewRouter function to create a router and connect handlers. It takes as input an intferface with database implementations.
func NewRouter(storageData *storage.Storage) *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.RequestID)
	router.Use(logmiddleware.LoggingMiddleware)

	router.Get("/api/ping", ping.Ping)
	router.Get("/api/tenders", tenders.GetTenderH(storageData))
	router.Get("/api/tenders/my", mytenders.GetMyTender(storageData))
	router.Get("/api/tenders/{tenderId}/status", tenderstatus.TenderStatus(storageData))
	router.Get("/api/bids/my", mybids.GetMyBids(storageData))
	router.Get("/api/bids/{tenderId}/list", tenderbids.TenderBidsH(storageData))
	router.Get("/api/bids/{bidId}/status", bidstatus.BidStatus(storageData))

	router.Put("/api/tenders/{tenderId}/status", changetenderstatus.ChangeTenderStatus(storageData))
	router.Put("/api/tenders/{tenderId}/rollback/{version}", rollbacktender.RollbackH(storageData))
	router.Put("/api/bids/{bidId}/status", changebidstatus.BidStatusChange(storageData))
	router.Put("/api/bids/{bidId}/rollback/{version}", rollbackbid.RollbackH(storageData))

	router.Post("/api/tenders/new", newtender.NewTenderH(storageData))
	router.Post("/api/bids/new", newbid.NewBidH(storageData))

	router.Patch("/api/tenders/{tenderId}/edit", edittender.EditTenderH(storageData))
	router.Patch("/api/bids/{bidId}/edit", editbid.EditBidH(storageData))

	router.NotFoundHandler()
	return router
}
