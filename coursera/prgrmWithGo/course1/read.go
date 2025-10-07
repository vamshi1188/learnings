package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type names struct {
	fname string
	lname string
}

func main() {

	// Open the file
	file, err := os.Open("names.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var people []names
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue // skip empty lines
		}

		parts := strings.Fields(line)
		if len(parts) < 2 {
			continue // skip invalid lines
		}

		person := names{fname: parts[0], lname: parts[1]}
		people = append(people, person)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// Print all names
	fmt.Println("\nNames in the file:")
	for _, p := range people {
		fmt.Println(p.fname, p.lname)
	}
}
