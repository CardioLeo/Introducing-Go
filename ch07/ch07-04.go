package main

import (
	"fmt"
)

type Circle struct {
	var x, y, r float64
}


type MultiShape struct {
	shapes []Shape
}

type Shape interface {
	area() float64
}

/*
func totalArea(circles ...Circle) float64 {
	var total float64
	for _, c := range circles {
		total += c.area()
	}
	return total
}
*/

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

func (m *MultiShape) area() float64 {
	var area float64
	for _, s := range m.shapes {
		area += s.area()
	}
	return area
}

func main() {
	// fmt.Println(totalArea(&c, &r))
	multiShape := MultiShape{
		shapes: []Shape{
			Circle{0, 0, 5},
			Rectangle{0, 0, 10, 15},
		},
	}
}

// without a coherent goal, I don't know how to shape this code into something ... coherent
