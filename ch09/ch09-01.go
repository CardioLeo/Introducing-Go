package main

import (
	"fmt"
//	"testing"
)

func TestAverage(/* t *testing.T */) {
	v := Average([]float64{1, 2})
	if v != 1.5 {
		fmt.Println("Error: Expected 1.5, got ", v)
	} else {
		fmt.Println("No errors found")
	}
}

func Average(xs []float64) float64 {
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
	fmt.Println(Average(xs))
	
	maximum(xs)
	minimum(xs)
	TestAverage()
}
