package router

import (
	"avitoTest/backend/interfaces/middleware/log"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// NewRouter function to create a router and connect handlers. It takes as input an intferface with database implementations.
func NewRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.RequestID)
	router.Use(logmiddleware.LoggingMiddleware)

	router.NotFoundHandler()
	return router
}
