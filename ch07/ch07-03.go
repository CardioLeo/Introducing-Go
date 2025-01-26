package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

func circleArea(c Circle) float64 {
	return math.Pi * c.r * c.r
}

func main() {
	// var c Circle
	// c := new(Circle)
	c1 := Circle{x: 0, y: 0, r: 5}
	// c2 := Circle(0, 0, 5) // gives error
	// c3 := &Circle(0, 0, 5)

	c1.x = 10
	c1.y = 5
	fmt.Println(c1.x, c1.y, c1.r)

	myC1Area := circleArea(c1)
	fmt.Println(myC1Area)
	// outputs: 78.53981633974483
}
