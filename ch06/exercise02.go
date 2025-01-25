package main

import (
	"fmt"
)

func half(input float64) (float64, bool) {
	var isEven bool
	if int(input) % 2 == 0 {
		isEven = true
	} else {
		isEven = false
	}
	outputNum := input / 2
	fmt.Println("For input", input, "when divided in half we get:", outputNum, "and it is", isEven, "that this number is even")
	return outputNum, isEven
}

func main() {
	half(3)
	half(20)
}
