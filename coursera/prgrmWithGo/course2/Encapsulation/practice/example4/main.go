package main

import "fmt"

type User struct {
	name      string
	birthyear int
	age       int
}

func (u *User) UserDetails(n string, by int) {
	if n == "" {
		fmt.Println("name cannot be empty!")
		return
	}
	u.name = n

	if by == 0 {
		fmt.Println("invalid birth year!")
		return
	}
	u.birthyear = by
}
func (u *User) GetCurrentAge(currentyear int) {

	u.age = currentyear - u.birthyear
}

func (u User) GetUserDetails() (string, int) {

	return u.name, u.age
}
func main() {
	var u User

	CurrentYear := 2025
	UserName := "vamshi"
	BirthYear := 2001

	u.UserDetails(UserName, BirthYear)
	u.GetCurrentAge(CurrentYear)
	fmt.Println(u.GetUserDetails())

}
