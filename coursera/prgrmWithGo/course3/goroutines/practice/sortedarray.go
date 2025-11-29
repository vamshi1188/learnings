package main

import (
	"fmt"
	"sort"
	"sync"
)

func SortArry() {
	// Step 1: Take user input
	fmt.Println("Enter integers separated by spaces:")
	var nums []int
	for {
		var n int
		_, err := fmt.Scan(&n)
		if err != nil {
			break
		}
		nums = append(nums, n)
	}

	if len(nums) == 0 {
		fmt.Println("No numbers provided.")
		return
	}

	// Step 2: Divide the array into 4 roughly equal parts
	n := len(nums)
	partSize := n / 4
	remainder := n % 4

	var parts [][]int
	start := 0
	for i := 0; i < 4; i++ {
		end := start + partSize
		if remainder > 0 {
			end++
			remainder--
		}
		if end > n {
			end = n
		}
		parts = append(parts, nums[start:end])
		start = end
	}

	// Step 3: Sort each part concurrently
	var wg sync.WaitGroup
	sortedParts := make([][]int, 4)

	for i, part := range parts {
		wg.Add(1)
		go func(i int, subarr []int) {
			defer wg.Done()
			fmt.Printf("Goroutine %d sorting subarray: %v\n", i+1, subarr)
			sort.Ints(subarr)
			sortedParts[i] = subarr
			fmt.Printf("Goroutine %d finished sorting: %v\n", i+1, subarr)
		}(i, append([]int(nil), part...)) // copy slice
	}

	wg.Wait()

	// Step 4: Merge the 4 sorted slices
	final := mergeAll(sortedParts)

	// Step 5: Print the final sorted array
	fmt.Println("Final sorted array:", final)
}

// mergeAll merges multiple sorted slices into one sorted slice.
func mergeAll(slices [][]int) []int {
	if len(slices) == 0 {
		return []int{}
	}
	result := slices[0]
	for i := 1; i < len(slices); i++ {
		result = mergeTwo(result, slices[i])
	}
	return result
}

// mergeTwo merges two sorted slices into one sorted slice.
func mergeTwo(a, b []int) []int {
	i, j := 0, 0
	var merged []int

	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			merged = append(merged, a[i])
			i++
		} else {
			merged = append(merged, b[j])
			j++
		}
	}
	merged = append(merged, a[i:]...)
	merged = append(merged, b[j:]...)
	return merged
}
