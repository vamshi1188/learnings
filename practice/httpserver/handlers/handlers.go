package handlers

import (
	"fmt"
	"net/http"
)

var names = make([]string, 0)

func HelloHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "hello i have starte an http server!")
}

func FormHandler(w http.ResponseWriter, r *http.Request) {

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

func Nameshandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, names)
}
