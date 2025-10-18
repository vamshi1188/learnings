package main

import (
	"bufio"
	"fmt"
	"os"

	"strconv"
)

func main() {

	sc := bufio.NewScanner(os.Stdin)

	fmt.Println("enter 10 numbers ")

	var numbers []int

	for i := 0; i < 10; i++ {

		sc.Scan()
		text := sc.Text()

		num, err := strconv.Atoi(text)

		if err != nil {

			fmt.Println("invalid numbewer ", err)
		}

		numbers = append(numbers, num)

	}
	fmt.Println("Numbers you entered:", numbers)

	BubbleSort(numbers)

	fmt.Println("Numbers sorted :", numbers)

}

func BubbleSort(a []int) {

	for j := 0; j < len(a)-1; j++ {

		for i := range a {

			if i < len(a)-1 {
				if a[i] > a[i+1] {

					Swap(a, i)

				}
			}
		}
	}
}
func Swap(a []int, i int) {
	a[i], a[i+1] = a[i+1], a[i]
}
