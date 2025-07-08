package routes

import (
	"femproject/internal/app"

	"github.com/go-chi/chi/v5"
)

func SetupRoutes(app *app.Application) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/health", app.Healthcheck)
	r.Get("/workouts/{id}", app.WorkoutHandler.GetHandlerbyId)
	r.Post("/workouts", app.WorkoutHandler.Handlercreateworkout)

	return r
}
