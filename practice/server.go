package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Welcme to my server..")
	http.HandleFunc("/count", Count)
	http.HandleFunc("/", Home)
	http.ListenAndServe("localhost:8080", nil)

}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "go to http://localhost:8080/count")
}

func Count(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "you are on count page")
}
