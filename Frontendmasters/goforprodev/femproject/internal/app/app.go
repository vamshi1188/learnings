package app

import (
	"database/sql"
	"femproject/internal/api"
	"femproject/internal/migrations"
	"femproject/internal/stores"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Application struct {
	Logger         log.Logger
	WorkoutHandler *api.WorkoutHandler
	DB             *sql.DB
}

func NewApplication() (*Application, error) {

	pgDB, err := stores.Open()
	if err != nil {

		return nil, err
	}

	err = stores.MigrateFS(pgDB, migrations.FS, ".")
	if err != nil {
		panic(err)
	}

	loggers := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	workoutStore := stores.NewPostgresWorkoutStore(pgDB)

	workouthanler := api.NewWorkoutHandler(workoutStore)

	app := &Application{
		Logger:         *loggers,
		WorkoutHandler: workouthanler,
		DB:             pgDB,
	}

	return app, nil
}

func (a *Application) Healthcheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "status is available\n")
}
