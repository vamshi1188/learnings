package main

import (
	"fmt"
	"sync"
)

func Sum(numbers []int, resulltchan chan<- int, wg *sync.WaitGroup) {
	wg.Done()

	sum := 0

	for _, num := range numbers {
		sum += num
	}
	resulltchan <- sum
}

func main() {

	var wg sync.WaitGroup

	resulltchan := make(chan int)

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	numberOfGoroutine := 4

	SliceSize := len(numbers) / numberOfGoroutine

	for i := 0; i < numberOfGoroutine; i++ {

		startIndex := 1 * SliceSize
		endIndex := (i + 1) * SliceSize

		if i == numberOfGoroutine-1 {
			endIndex = len(numbers)
		}
		wg.Add(1)
		go Sum(numbers[startIndex:endIndex], resulltchan, &wg)
	}

	var collectingwg sync.WaitGroup

	collectingwg.Add(1)
	go func() {
		defer collectingwg.Done()
		totalsum := 0
		for i := 0; i < numberOfGoroutine; i++ {

			partialSum := <-resulltchan

			totalsum += partialSum
		}
		fmt.Printf("sum of total numbers is %d", totalsum)
	}()

	wg.Wait()
	close(resulltchan)
	collectingwg.Wait()
}
