package main

import (
	"fmt"
	"sort"
	"strconv"
)

func Slices() {
	// Create an empty slice of integers with initial capacity 3
	nums := make([]int, 0, 3)

	fmt.Println("Enter integers to add to the slice. Enter 'X' to quit.")

	for {
		// Prompt the user
		fmt.Print("Enter an integer (or 'X' to quit): ")
		var input string
		fmt.Scan(&input)

		if input == "X" || input == "x" {
			fmt.Println("Exiting...")
			break
		}
		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input, please enter an integer or 'X'.")
			continue
		}
		nums = append(nums, num)
		sort.Ints(nums)
		fmt.Println("Sorted slice:", nums)
	}
}
