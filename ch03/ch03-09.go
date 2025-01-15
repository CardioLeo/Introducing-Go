package main

import (
	"fmt"
)

func main() {
	const x string = "Hello, World"
	fmt.Println(x)

	var (
		a = 5
		b = 10
		c = 15
	)
	fmt.Println(a, b, c, a+b+c)
}
