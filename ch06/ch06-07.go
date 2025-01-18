package main

import (
	"fmt"
)

func factorial(x int) int {
	if x == 0 {
		return 1
	}
	fmt.Println(x)
	return x * factorial(x-1)
}

func ask_for_input() int {
	fmt.Println("What number would you like to search for the factorial for?")
	input, ok := fmt.Scanln()
	if ok == nil {
		panic("ok is equal to nil")
	}
	return input
}

func main() {
	input := ask_for_input()
	fmt.Println(factorial(input))
}
