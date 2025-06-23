package router

import (
	"app-destination/database"
	"app-destination/handler"
	"app-destination/middleware"
	"app-destination/repository"
	"app-destination/service"
	"log"

	"github.com/go-chi/chi"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Could not initialize database: %v", err)
	}
	repo := repository.NewDestinationRepository(db)
	srv := service.NewDestinationService(repo)
	h := handler.NewDestinationHandler(srv)

	r.Use(middleware.Logger)
	r.Use(middleware.BasicAuth)

	r.Get("/api/destinations", h.GetAllDestinations)
	r.Post("/api/destinations", h.CreateDestination)

	return r

}