package main

import "fmt"

type User struct {
	name string
}

func (u *User) SetName(n string) {

	if n == "" {
		fmt.Println("name cannot be empty!")
		return
	}
	u.name = n
}

func (u User) GetName() string {

	return u.name
}

func main() {

	var u User

	name := "vamshi"
	name2 := ""
	u.SetName(name2)
	fmt.Println(u.GetName())

	u.SetName(name)
	fmt.Println(u.GetName())
}
