package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/ikhwan-satrio/auth-golang/app/handlers"
	"github.com/ikhwan-satrio/auth-golang/app/middlewares"
)

func SetupRoutes(app *chi.Mux, authHandler *handlers.AuthHandler, userHandler *handlers.UserHandler) {
	app.Use(middleware.Logger)
	app.Use(middleware.Recoverer)

	app.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("auth-golang service is running"))
	})

	app.Route("/auth", func(r chi.Router) {
		r.Post("/register", authHandler.Register)
		r.Post("/login", authHandler.Login)
	})

	app.With(middlewares.AuthMiddleware).Group(func(r chi.Router) {
		r.Get("/users", userHandler.GetUsers)
	})
}
