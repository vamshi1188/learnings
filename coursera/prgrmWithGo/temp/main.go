package main

import (
	"fmt"
)

func main() {

	var word string

	count := 0
	count2 := 0
	count3 := 0

	fmt.Println(" Enter a string i will check for 'i,a,n'")
	fmt.Scan(&word)

	for _, r := range word {

		switch r {
		case 'i':
			count = 1
		case 'a':
			count2 = 1
		case 'n':
			count3 = 1
		}
	}

	if count == 1 && count2 == 1 && count3 == 1 {
		fmt.Println("Found")
	} else {
		fmt.Println("Notfound")
	}
}
