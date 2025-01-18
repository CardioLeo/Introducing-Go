package main

import (
	"fmt"
)

func convert_to_Celsius(input float64) float64{
	var C float64
	C = input - 32
	C = C * 5 / 9
	return C
	// I tested it, and this works the same as what they suggested
	// they have it written out as x := (input - 32) * 5/9
	// there's is probably better
}

func main() {
	fmt.Println("Enter a number to convert from Fahrenheit to Celsius: ")
	var input float64
	fmt.Scanf("%f", &input)
	output := convert_to_Celsius(input)
	fmt.Println(output)
}
