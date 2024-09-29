package ping

import "github.com/go-chi/chi/v5"

func PingController(router *chi.Mux) {
	router.Get("/ping", ping)
}
