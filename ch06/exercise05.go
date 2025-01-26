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

func recursive_fib(times int) int {
	switch times {
		case 0:
			return 0
		case 1:
			return 1
		default:
			return recursive_fib(times - 1) + recursive_fib(times - 2)
			// temp := recursive_fib(times - 1) + recursive_fib(times - 2)
			// fmt.Println(temp)
			// return temp

			// these commented out lines show just how efficient my other
			// solution to this problem was - even though I solved it that
			// way initailly due to anxiety from the way the recursive 
			// one was working
	}
}

func main() {
	fib(5)
	fib(13)

	fmt.Println(recursive_fib(13))
}
