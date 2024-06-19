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
		r.Get("/feed", db.HandleGetFeeds)
		r.Get("/follow", db.AuthUser(db.HandleGetFollows))
		r.Get("/posts", db.AuthUser(db.HandleGetUserPosts))

		r.Post("/user", db.HandleCreateUser)
		r.Post("/feed", db.AuthUser(db.HandleCreateFeed))
		r.Post("/follow", db.AuthUser(db.HandleFollowFeed))

		r.Delete("/follow/{followId}", db.AuthUser(db.HandleDeleteFollow))
	})

	return router
}