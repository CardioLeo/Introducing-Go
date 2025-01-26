package main

import (
	"fmt"
	"math"
)

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func rectangleArea(x1, y1, x2, y2 float64) float64 {
	l := distance(x1, y1, x1, y2) 
	w := distance(x1, y1, x2, y1) // if you're not careful you'll miss this
	return l * w
}

func circleArea(r float64) float64 {
	return math.Pi * r * r // but why are x and y passed in here?
				// yeah, it returns a compiler error
}

func main() {
	var rx1, ry1 float64 = 0.0, 0.0
	var rx2, ry2 float64 = 10.0, 10.0
	var cr float64 = 5.0

	fmt.Println(rectangleArea(rx1, ry1, rx2, ry2))
	fmt.Println(circleArea(cr))
}

// this changes the code in the book enough to make it compile
// do they deliberately make the code such that it won't compile
// so that you'll be encouraged (if not forced) to test it out and learn yourself?
