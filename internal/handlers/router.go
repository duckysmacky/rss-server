package handlers

import (
	"github.com/go-chi/chi"
	chimiddleware "github.com/go-chi/chi/middleware"
)

func NewRouter() *chi.Mux {
	var router = chi.NewMux()
	router.Use(chimiddleware.StripSlashes)

	router.Route("/api", func(api chi.Router) {
		api.Get("/status", responseStatus)
		api.Get("/err", responseBadRequest)
	})

	return router
}