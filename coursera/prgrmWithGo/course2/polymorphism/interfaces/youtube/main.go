package main

import "fmt"

type shape interface {
	ShapeName() string
}
type perimeter interface {
	Area() float64
}

type Geometry interface {
	shape
	perimeter
}

type Rectangle struct {
	width  float64
	height float64
	name   string
}

func (s Rectangle) Area() float64 {
	return s.width * s.height
}

func (d Rectangle) ShapeName() string {
	return d.name
}

func main() {
	var g Geometry

	g = Rectangle{width: 5, height: 3, name: "Rectangle"}
	fmt.Println(g.Area(), g.ShapeName())
}
