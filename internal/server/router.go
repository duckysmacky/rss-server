package server

import (
	"github.com/duckysmacky/rss-server/internal/api"
	"github.com/duckysmacky/rss-server/internal/handlers"
	"github.com/go-chi/chi"
	chimiddleware "github.com/go-chi/chi/middleware"
)

func newRouter() *chi.Mux {
	var db = handlers.Database
	var router = chi.NewMux()
	router.Use(chimiddleware.StripSlashes)

	router.Route("/api", func(r chi.Router) {
		r.Get("/status", api.ResponseStatus)
		r.Get("/user", db.AuthUser(db.HandleGetUser))
		r.Get("/feeds", db.HandleGetFeeds)
		r.Get("/follows", db.AuthUser(db.HandleGetFollows))

		r.Post("/user", db.HandleCreateUser)
		r.Post("/feed", db.AuthUser(db.HandleCreateFeed))
		r.Post("/follows", db.AuthUser(db.HandleFollowFeed))

		r.Delete("/follows/{followId}", db.AuthUser(db.HandleDeleteFollow))
	})

	return router
}