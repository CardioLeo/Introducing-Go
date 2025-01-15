package main

import (
	"fmt"
)

func early_func() {
	x := [4]float64{
		98,
		93,
		77,
		82,
		// 83,
	}
	var total float64 = 0
	for _, value := range x {
		total += value
	}
	fmt.Println(total / float64(len(x)))
}

func main() {
	early_func()
	// var x []float64 // this is a slice!
			// because of the missing array length!
			// really good to know!
}
