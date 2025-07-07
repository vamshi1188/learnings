package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type Application struct {
	Logger log.Logger
}

func NewApplication() (*Application, error) {

	loggers := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &Application{
		Logger: *loggers,
	}

	return app, nil
}

func (a *Application) Healthcheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "status is available\n")
}
