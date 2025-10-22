package main

import "fmt"

type User1 struct {
	name string
	age  int
}

type User2 struct {
	name string
	age  int
}

type User3 struct {
	name string
	age  int
}

func (u *User1) SetNameAndAge1(n string, a int) {

	if n == "" {
		fmt.Println("name cannot be empty")
		return
	}
	u.name = n
	if a == 0 {
		fmt.Println("inavlid age ")
		return
	}
	u.age = a
}

func (u *User2) SetNameAndAge2(n string, a int) {

	if n == "" {
		fmt.Println("name cannot be empty")
		return
	}
	u.name = n
	if a == 0 {
		fmt.Println("inavlid age ")
		return
	}
	u.age = a
}
func (u *User3) SetNameAndAge3(n string, a int) {

	if n == "" {
		fmt.Println("name cannot be empty")
		return
	}
	u.name = n
	if a == 0 {
		fmt.Println("inavlid age ")
		return
	}
	u.age = a
}
func (u User1) GetNameAndAge1() (string, int) {

	return u.name, u.age
}
func (u User2) GetNameAndAge2() (string, int) {

	return u.name, u.age
}
func (u User3) GetNameAndAge3() (string, int) {

	return u.name, u.age
}
func main() {
	var u1 User1
	var u2 User2
	var u3 User3

	u1.name = "vamshi"
	u1.age = 24
	u2.name = "sai"
	u2.age = 22
	u3.name = "saivamshi"
	u3.age = 20

	u1.SetNameAndAge1(u1.name, u1.age)
	u2.SetNameAndAge2(u2.name, u2.age)
	u3.SetNameAndAge3(u3.name, u3.age)

	fmt.Println(u1.GetNameAndAge1())
	fmt.Println(u2.GetNameAndAge2())
	fmt.Println(u3.GetNameAndAge3())
}
