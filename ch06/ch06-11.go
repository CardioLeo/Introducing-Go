package main

import (
	"fmt"
)

func zero(x int) {
	x = 0
}

func zero2(xPtr *int) {
	*xPtr = 0
}

func main() {
	x := 5
	zero(x)
	fmt.Println(x) // x is still 5 // <-- Doxsey's comment
	zero(&x)
	fmt.Println(x) // x is 0 // <-- Doxsey's comment
}
