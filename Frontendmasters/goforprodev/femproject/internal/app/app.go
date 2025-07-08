package app

import (
	"femproject/internal/api"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Application struct {
	Logger         log.Logger
	WorkoutHandler *api.WorkoutHandler
}

func NewApplication() (*Application, error) {

	loggers := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	workouthanler := api.NewWorkoutHandler()

	app := &Application{
		Logger:         *loggers,
		WorkoutHandler: workouthanler,
	}

	return app, nil
}

func (a *Application) Healthcheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "status is available\n")
}
