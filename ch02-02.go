package main

import "fmt"

func main() {
	my_string := "Oh My Gosh"
	fmt.Println("My String is:", my_string)
	my_string_length := len(my_string)
	fmt.Println("My String Length is:", my_string_length)
	my_string_example := "my string example"[12]
	fmt.Println("My String Example, printed with fmt.Println:", my_string_example)
	fmt.Printf("My String Example, printed with fmt.Printf:", my_string_example, "\n")
	fmt.Println("My " + "Strings" + " Three")
}
