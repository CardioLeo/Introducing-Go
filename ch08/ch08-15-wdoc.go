package main

import (
	"fmt"
	"flag"
	"math/rand"
)

func main() {
	// Define flags
	maxp := flag.Int("max", 6, "the max value")
	// parse
	flag.Parse()
	// Generate a number between 0 and max
	fmt.Println(rand.Intn(*maxp))
}
