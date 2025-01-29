package main

import (
	"fmt"
)

func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

func maximum(xs []float64) float64 {
	temp := 0.0
	for _, v := range xs {
		if v > temp {
			temp = v
			fmt.Println("the new maximum is", temp)
		}
	}
	return temp
}

func minimum(xs []float64) float64 {
	temp := 100.0
	for _, v := range xs {
		if v < temp {
			temp = v
			fmt.Println("the new minimum is", temp)
		}
	}
	return temp
}

func main() {
	xs := []float64{64, 83, 77, 82, 84, 54, 98}
	fmt.Println(average(xs))
	
	maximum(xs)
	minimum(xs)
}
