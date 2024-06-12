package routers

import "github.com/go-chi/chi"

func NewRouter() *chi.Mux {
	var mux = chi.NewMux()

	return mux
}