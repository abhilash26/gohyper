package router

import (
	"github.com/abhilash26/gohyper/internal/handler"
	"github.com/abhilash26/gohyper/internal/middleware"
	"github.com/go-chi/chi/v5"
	cmiddleware "github.com/go-chi/chi/v5/middleware"
)

func SetupWebRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(cmiddleware.Logger)

	/* If using nginx or apache, disable below middleware and
	   pass static resources directly through them */
	router.Use(middleware.FileServer("/static", "public"))

	// Application Routes here
	router.Get("/", handler.IndexHandler)
	router.Get("/counter-subtract", handler.CounterSubtract)
	router.Get("/counter-add", handler.CounterAdd)

	return router
}
