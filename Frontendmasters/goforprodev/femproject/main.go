package main

import (
	"femproject/internal/app"
	"femproject/internal/routes"
	"flag"
	"fmt"
	"net/http"
	"time"
)

func main() {

	var port int

	flag.IntVar(&port, "port", 8080, "backend server port")
	flag.Parse()

	app, err := app.NewApplication()

	if err != nil {

		panic(err)
	}

	app.Logger.Printf("we are runnig on %d\n", port)

	r := routes.SetupRoutes(app)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	err = server.ListenAndServe()
	if err != nil {

		app.Logger.Fatal(err)
	}
}
