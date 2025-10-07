package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func stringg() {
	fmt.Println("Enter a string:")

	// Read the entire line (with spaces if any)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	// Trim spaces and newline, convert to lowercase
	input = strings.TrimSpace(strings.ToLower(input))

	// Check conditions
	if strings.HasPrefix(input, "i") &&
		strings.HasSuffix(input, "n") &&
		strings.Contains(input, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
