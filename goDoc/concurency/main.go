package main

import (
	"fmt"
	"sync"
	"time"
)

func PrintOne(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i <= 50; i++ {

		fmt.Print("1")

		time.Sleep(10 * time.Millisecond)

	}
}
func PrintTwo(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i <= 50; i++ {

		fmt.Print("2")

		time.Sleep(10 * time.Millisecond)

	}
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	go PrintOne(&wg)

	go PrintTwo(&wg)
	wg.Wait()
}
