package main

import (
	"fmt"
)

func main() {
	defer func() {
		str := recover() // prints nothing I can see?
		str2 := "did this work?" // works
		fmt.Println(str, str2)
	}()
	panic("PANIC")
}
