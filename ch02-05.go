package main

import "fmt"

func main() {
	fmt.Println("They asked me to print out this value to figure out what it is!")
	fmt.Println((true && false) || (false && true) || !(false && false))
}
