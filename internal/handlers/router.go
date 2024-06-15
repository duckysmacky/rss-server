package handlers

import (
	"github.com/duckysmacky/rss-server/internal/api"
	"github.com/go-chi/chi"
	chimiddleware "github.com/go-chi/chi/middleware"
)

func newRouter(db Database) *chi.Mux {
	var router = chi.NewMux()
	router.Use(chimiddleware.StripSlashes)

	router.Route("/api", func(r chi.Router) {
		r.Get("/status", api.ResponseStatus)
		r.Post("/user", db.handleCreateUser)
	})

	return router
}