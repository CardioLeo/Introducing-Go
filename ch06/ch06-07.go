package main

import (
	"fmt"
)

func factorial(x int64) int64 {
	if x == 0 {
		return 1
	}
	fmt.Println(x)
	return x * factorial(x-1)
}

func ask_for_input() int64 {
	fmt.Println("What number would you like to search for the factorial for?")
	var input int64
	fmt.Scanf("%f", &input)
	return input
}

// this doesn't really run like I want it to
// 1. it won't take more than a single digit as input
// 2. Won't ask for input at all if I use Scanln (scanf > scanln)
// 3. it doesn't seem to be the full logic for a factorial?
// (I just based what I wrote on what the book provided)
// 

func main() {
	input := ask_for_input()
	fmt.Println(factorial(input))
}
