package main

import (
	"fmt"
	"net/http"
)

func main() {

	fmt.Println("server started on port 8080")

	http.HandleFunc("/", HomePage)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("error in setup the server!")
		return
	}

}

func HomePage(w http.ResponseWriter, r *http.Request) {
	res, err := http.Get("https://railradar.in/api/v1/search/stations")
	if err != nil {
		fmt.Fprintln(w, err)
	}
	res.Body.Close()
	fmt.Fprintln(w, res)
}
