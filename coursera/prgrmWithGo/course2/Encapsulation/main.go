package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (e Animal) Eat() {

	fmt.Println(e.food)
}
func (e Animal) Move() {

	fmt.Println(e.locomotion)
}
func (e Animal) Speak() {

	fmt.Println(e.noise)
}
func main() {

	cow := Animal{food: "grass", locomotion: "walk", noise: "moo"}
	bird := Animal{food: "worms", locomotion: "fly", noise: "peep"}
	snake := Animal{food: "mice", locomotion: "slither", noise: "hsss"}

	sc := bufio.NewScanner(os.Stdin)

	for {

		fmt.Print(">")
		sc.Scan()
		line := sc.Text()

		field := strings.Fields(line)
		animal := field[0]
		action := field[1]

		switch animal {
		case "cow":
			switch action {
			case "eat":
				cow.Eat()
			case "move":
				cow.Move()
			case "speak":
				cow.Speak()
			}
		case "bird":
			switch action {
			case "eat":
				bird.Eat()
			case "move":
				bird.Move()
			case "speak":
				bird.Speak()
			}
		case "snake":
			switch action {
			case "eat":
				snake.Eat()
			case "move":
				snake.Move()
			case "speak":
				snake.Speak()
			}

		}

	}

}
