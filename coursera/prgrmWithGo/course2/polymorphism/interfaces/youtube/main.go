package main

import "fmt"

type Shape interface {
	Area() float64
}
type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Area() float64 {

	return r.width * r.height

}

func Calculation(s Shape) float64 {
	return s.Area()
}

func main() {
	rect := Rectangle{width: 5, height: 4}

	fmt.Println(Calculation(rect))
}
