package main

import (
	"fmt"
)

// GenDisplaceFn returns a function that calculates displacement based on given a, vo, and so
func GenDisplaceFn(a, vo, so float64) func(float64) float64 {
	return func(t float64) float64 {
		// Formula: s = ½ a t² + vo*t + so
		return 0.5*a*t*t + vo*t + so
	}
}

func main() {
	var a, vo, so, t float64

	// Get user input
	fmt.Print("Enter acceleration (a): ")
	fmt.Scan(&a)

	fmt.Print("Enter initial velocity (vo): ")
	fmt.Scan(&vo)

	fmt.Print("Enter initial displacement (so): ")
	fmt.Scan(&so)

	// Generate the displacement function
	fn := GenDisplaceFn(a, vo, so)

	// Ask for time
	fmt.Print("Enter time (t): ")
	fmt.Scan(&t)

	// Compute and print displacement
	displacement := fn(t)
	fmt.Printf("Displacement after %.2f seconds is: %.2f\n", t, displacement)
}
