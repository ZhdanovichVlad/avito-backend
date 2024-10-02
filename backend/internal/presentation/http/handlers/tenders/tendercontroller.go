package tender

import (
	tenderapplication "avitoTest/backend/internal/application/tender"
	"github.com/go-chi/chi/v5"
)

// Constants for paths
const (
	tenderBasePath      = "/api/tenders"
	newTenderPath       = "/new"
	tenderStatusPath    = "/{tenderId}/status"
	myTendersPath       = "/my"
	PutTenderStatusPath = "/{tenderId}/status"
)

type TenderController struct{}

func AddTenderController(router *chi.Mux, storageData tenderapplication.Application) {
	router.Route(tenderBasePath, func(r chi.Router) {

		r.Post(newTenderPath, CreateNewTender(storageData))
		r.Get(myTendersPath, GetMyTendersH(storageData))

		r.Get("/", GetTendersH(storageData))
		r.Get(tenderStatusPath, GetTenderStatusH(storageData))

		r.Put(PutTenderStatusPath, PutTenderStatus(storageData))
	})
}
