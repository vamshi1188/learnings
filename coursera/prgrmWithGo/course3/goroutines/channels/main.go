package main

import "fmt"

func main() {

	ch := make(chan int)

	ch <- 3

	num := <-ch
	fmt.Println(num)
}

func Communication(a int, b int, c chan int) int {

	re
}
