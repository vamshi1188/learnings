package main

import "fmt"

func Anonymous() {
	func(name string) {
		fmt.Println("hey", name)
	}("saivamshi")
}
