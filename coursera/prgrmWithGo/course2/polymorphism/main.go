package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Interface for all animals
type Animal interface {
	Eat()
	Move()
	Speak()
}

// Cow, Bird, Snake types
type Cow struct{}
type Bird struct{}
type Snake struct{}

// Implement methods for Cow
func (c Cow) Eat()   { fmt.Println("grass") }
func (c Cow) Move()  { fmt.Println("walk") }
func (c Cow) Speak() { fmt.Println("moo") }

// Implement methods for Bird
func (b Bird) Eat()   { fmt.Println("worms") }
func (b Bird) Move()  { fmt.Println("fly") }
func (b Bird) Speak() { fmt.Println("peep") }

// Implement methods for Snake
func (s Snake) Eat()   { fmt.Println("mice") }
func (s Snake) Move()  { fmt.Println("slither") }
func (s Snake) Speak() { fmt.Println("hsss") }

func main() {
	animals := make(map[string]Animal) // name → animal
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ") // prompt

		scanner.Scan()
		input := scanner.Text()
		parts := strings.Fields(input)

		if len(parts) != 3 {
			fmt.Println("Invalid command. Use 'newanimal' or 'query'.")
			continue
		}

		command, name, info := parts[0], parts[1], parts[2]

		switch command {
		case "newanimal":
			switch info {
			case "cow":
				animals[name] = Cow{}
			case "bird":
				animals[name] = Bird{}
			case "snake":
				animals[name] = Snake{}
			default:
				fmt.Println("Unknown animal type. Use cow, bird, or snake.")
				continue
			}
			fmt.Println("Created it!")

		case "query":
			animal, found := animals[name]
			if !found {
				fmt.Println("Animal not found!")
				continue
			}

			switch info {
			case "eat":
				animal.Eat()
			case "move":
				animal.Move()
			case "speak":
				animal.Speak()
			default:
				fmt.Println("Unknown query type. Use eat, move, or speak.")
			}

		default:
			fmt.Println("Invalid command. Use 'newanimal' or 'query'.")
		}
	}
}
