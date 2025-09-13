package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", homehandler)
	http.ListenAndServe(":8080", nil)
}

func homehandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello!")
}
