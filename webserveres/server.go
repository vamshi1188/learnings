package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
)

type User struct {
	Name            string
	car             string
	chargercapacity string
}

func main() {
	fmt.Println("welcome to web  server no 1")

	http.HandleFunc("/", Homepage)
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("error in opening server on port 8080")
	}

}

func Homepage(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "woke up servewr 1")
	var u User
	fmt.Println(UserDetails(&u))
}

func UserDetails(p *User) (string, string, string) {
	sc := bufio.NewScanner(os.Stdin)
	fmt.Println("what is your name ? ")
	sc.Scan()
	p.Name = sc.Text()
	fmt.Println("what is your car  model?")
	sc.Scan()
	p.car = sc.Text()
	fmt.Println("how much capacity cahrger you want?")
	sc.Scan()
	p.chargercapacity = sc.Text()
	return p.Name, p.car, p.chargercapacity
}
