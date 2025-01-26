package main

import (
	"fmt"
)

func swap(x *int, y *int) {
	temp := *x
	*x = *y
	*y = temp
}

func main() {
	x := 1
	y := 2
	
	fmt.Println("Before calling swap, the value of x is:", x, "and the value of y is:", y)

	swap(&x, &y)
	
	fmt.Println("After calling swap, the value of x is:", x, "and the value of y is:", y)
}
