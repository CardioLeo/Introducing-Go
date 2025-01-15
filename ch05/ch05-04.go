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
	x := make([]float64, 5)
	y := make([]float64, 5, 10)
	arr := [5]float64{1, 2, 3, 4, 5}
	z := arr[0:5]
	// a := [3]float64{x, y, z,}
	var second_total float64
	for _, value := range x {
		second_total += value
	}
	fmt.Println(second_total / float64(len(x)))
	var third_total float64
	for _, value := range y {
		third_total += value
	}
	fmt.Println(third_total / float64(len(y)))
	var fourth_total float64
	for _, value := range z {
		fourth_total += value
	}
	fmt.Println(fourth_total / float64(len(z)))
}
