package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"github.com/avukadin/goapi/internal/middleware"
)

func Handler(r *chi.Mux) {
	// capital 'h' in handler above makes this globally available rather than just in this file
	r.Use(chimiddle.StripSlashes)

	r.Route("/account"), func(router, chi.Router) {
		router.Use(middleware.Authorization)

		router.Get("/coins", GetCoinBalance)
	}
}