package main

import "fmt"

type User struct {
	name string
}

func main() {

	var u User

	u.name = "vamshi"

	fmt.Println(u.name)

	// var user User
	// name := "vamshi"

	// user.SetName(name)
	// fmt.Println(user.GetName())
}

// func (u *User) SetName(n string) {
// 	u.name = n
// }

// func (u User) GetName() string {
// 	return u.name
// }
