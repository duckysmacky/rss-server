package routers

import (
	"github.com/duckysmacky/rss-server/internal/handlers"
	"github.com/go-chi/chi"
	chimiddleware "github.com/go-chi/chi/middleware"
)

func NewRouter() *chi.Mux {
	var router = chi.NewMux()
	router.Use(chimiddleware.StripSlashes)

	router.Route("/api", func(api chi.Router) {
		api.HandleFunc("/status", handlers.ResponseStatus)
	})

	return router
}