package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/ikhwan-satrio/auth-golang/app/config"
	"github.com/ikhwan-satrio/auth-golang/app/db"
	"github.com/ikhwan-satrio/auth-golang/app/handlers"
	"github.com/ikhwan-satrio/auth-golang/app/routes"
	"github.com/ikhwan-satrio/auth-golang/app/services/auth"
)

func main() {
	// 1. Config
	cfg := config.LoadConfig()

	// 2. Database
	database := &db.DBService{
		Config: cfg.DB,
	}
	dbConn, _, err := database.CreateDB()
	if err != nil {
		log.Fatalf("cannot initialize database: %v", err)
	}

	// 3. Services
	authService := auth.NewAuthService(dbConn)

	// 4. Handlers
	authHandler := handlers.NewAuthHandler(authService)
	userHandler := handlers.NewUserHandler(dbConn)

	// 5. Routes
	app := chi.NewRouter()
	routes.SetupRoutes(app, authHandler, userHandler)

	// 6. Start Server
	log.Printf("listening on http://localhost:%s", cfg.HTTP.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", cfg.HTTP.Port), app); err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}
