package routes

import (
	"femproject/internal/app"

	"github.com/go-chi/chi"
)

func SetupRoutes(app *app.Application) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/health", app.Healthcheck)

	return r
}
