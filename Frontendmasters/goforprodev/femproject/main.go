package main

import (
	"femproject/internal/app"
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

	http.HandleFunc("/healthcheck", Healthcheck)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	err = server.ListenAndServe()
	if err != nil {

		app.Logger.Fatal(err)
	}
}

func Healthcheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "status is available\n")
}
