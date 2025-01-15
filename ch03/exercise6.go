package main

import (
	"fmt"
)

func convert_feet_to_meters(feet float64) float64 {
	meters := feet * 0.3048
	return meters
}

func main() {
	fmt.Println("Enter the measurement in feet you want to convert to meters: ")
	var feet float64
	fmt.Scanf("%f", &feet)
	meters := convert_feet_to_meters(feet)
	fmt.Println(meters)
}
