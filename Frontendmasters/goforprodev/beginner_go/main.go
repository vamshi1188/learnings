package main

import "fmt"

func main() {
	const (
		s int = iota + 1
		d
		h
		jsd
		k
	)

	array := []int{2}
	fmt.Println(s, d, h, k, jsd)
	println(array)
	fmt.Println(array)

	capitalcities := map[string]string{

		"India": "delhi",
	}

	capital, exists := capitalcities["India"]

	if exists {
		println("yes it exists", capital)
	} else {
		fmt.Println("capital doesn't exist")
	}

	fmt.Printf("the capital of india is %v\n", capitalcities["India"])
}
