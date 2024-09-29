package tenders

import (
	"github.com/go-chi/chi/v5"

	tenderApplication "avitoTest/backend/application/tender"
)

// Constants for paths
const (
	tenderBasePath      = "/api/tenders"
	newTenderPath       = "/new"
	tenderStatusPath    = "/{tenderId}/status"
	myTendersPath       = "/my"
	PutTenderStatusPath = "/{tenderId}/status"
)

func AddTenderController(router *chi.Mux, storageData tenderApplication.NewTenderApplication) {
	router.Route(tenderBasePath, func(r chi.Router) {

		r.Post(newTenderPath, CreateNewTender(storageData))
		r.Get(myTendersPath, GetMyTendersH(storageData))

		r.Get("/", GetTendersH(storageData))
		r.Get(tenderStatusPath, GetTenderStatusH(storageData))

		r.Put(PutTenderStatusPath, PutTenderStatus(storageData))
	})
}
