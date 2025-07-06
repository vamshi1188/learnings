package app

import (
	"log"
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
