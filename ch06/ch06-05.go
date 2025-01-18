package main

import (
	"fmt"
)

func main() {
	add := func(x, y int) int {
		return x + y
	}
	fmt.Println(add(1,1))

	z := 0
	increment := func() int {
		z++
		return z
	}
	fmt.Println(increment())
	fmt.Println(increment())
}
