package main

import (
	"fmt"
)

var x string = "Hello, World"

func main() {
	fmt.Println(x)
}

func f() {
	fmt.Println(x)
}	// weird, because it doesn't call or use this function,
	// and it isn't exported
