package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		fmt.Println("goroutine")

	}()

	fmt.Println("main goroutine")

	wg.Wait()
}
