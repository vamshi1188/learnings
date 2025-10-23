package main

import "fmt"

type User struct {
	name string
	age  int
}

func (u *User) SetAge(n int) {
	if n < 1 || n > 120 {

		fmt.Println(" age has to be in between 1 to 120")

		return
	}

	u.age = n
}

func (u User) GetAge() int {
	return u.age
}

func main() {
	var u User
	Age := 0
	Age2 := 100
	Age3 := 121

	u.SetAge(Age)
	fmt.Println(u.GetAge())
	u.SetAge(Age2)
	fmt.Println(u.GetAge())
	u.SetAge(Age3)
	fmt.Println(u.GetAge())

}
