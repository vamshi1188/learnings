package main

import (
	"fmt"
	"time"
)

func main() {
	var x int

	go func() {
		for i := 0; i < 1000; i++ {
			x++ // 👈 concurrent write
		}
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			fmt.Print(x) // 👈 concurrent read
		}
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("\nDone")
}

// Both goroutines run for a while (1000 times each).

// One writes (x++), the other reads (fmt.Print(x)).

// They truly overlap in time.

// The -race detector catches the concurrent access to the same memory location.
// ==================
// WARNING: DATA RACE
// Read at 0x000...
// Write at 0x000...
// ...
// Found 1 data race(s)
// exit status 66
