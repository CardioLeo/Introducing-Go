package main

import (
	"fmt"
)

func fib(times int) {
	a, b, c := 1, 0, 0
	for i := 0; i < times; i++ {
		c = a + b
		a = b
		b = c
		fmt.Println(c)
	}
}

func main() {
	// fib(5)
	fib(13)
}
