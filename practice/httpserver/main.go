package main

import (
	"httpserver/root"
	"net/http"
	"net/url"
)

func main() {

	app := root.RouteMux()

	http.ListenAndServe(":8080", app)

}
func query() url.Values {

}
