package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"

	"github.com/st-ember/chessbackend/internal/handlers/auth"
	// "github.com/st-ember/chessbackend/internal/middleware"
)

func Handler(r *chi.Mux) {
	r.Use(chimiddle.StripSlashes)

	r.Route("/user", func(router chi.Router) {
		// router.Use(middleware.Authorization)

		router.Post("/login", auth.Login)
		router.Post("signup", auth.Signup)
	})
}
