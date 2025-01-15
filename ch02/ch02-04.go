package main

import (
	"fmt"
)

func main() {
	fmt.Println("the given number is a multiple of...")
	fmt.Println("32132 * 42452")
	fmt.Println(32132 * 42452)
}

// got an error saying math was not used; lets see if we really need that at all
// works nicely
// the output is here:
// the given number is a multiple of...
// 32132 * 42452
// 1364067664
