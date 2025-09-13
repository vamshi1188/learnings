package main

import (
	"fmt"
	"net/http"
)

func main() {

	var name string
	var decision string

	fmt.Println("hello welcome to the serverworlds!")
	fmt.Print("what is your name! :")
	fmt.Scan(&name)
	fmt.Printf("nice to meet you %v!\n", name)
	fmt.Print("do you want me to start server? :")
	fmt.Scan(&decision)

	if decision == "yes" || decision == "YES" {

		http.HandleFunc("/", Homehandler)
		http.HandleFunc("/about", AboutHandler)

		fmt.Printf("server started at port 8080!")

		http.ListenAndServe(":8080", nil)

	}

	fmt.Printf("bye %v am ending the session\n", name)

}

func Homehandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Fprintf(w, "hello,  %s\n", name)
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is about page")
}
