package root

import (
	"httpserver/handlers"
	"net/http"
)

func RouteMux() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.HelloHandler)
	mux.HandleFunc("/form", handlers.FormHandler)
	mux.HandleFunc("/names", handlers.Nameshandler)
	return mux
}
