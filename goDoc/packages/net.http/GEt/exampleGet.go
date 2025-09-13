package main

import (
	"fmt"
	"net/http"
)

var names = make([]string, 0)

func main() {

	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/names", nameshandler)

	http.ListenAndServe(":8080", nil)

}

func helloHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "hello i have starte an http server!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		r.ParseForm()

		name := r.FormValue("name")

		names = append(names, name)
		// fmt.Fprintf(w, "hello %s ", name)
		fmt.Println(names)

	}

	html := `
	<form method="POST">
			Name: <input name="name">
			<input type="submit">
		</form>
	`
	w.Header().Set("content-Type", "text/html")

	fmt.Fprintln(w, html)
}

func nameshandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, names)
}
