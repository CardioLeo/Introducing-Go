package main

import (
	"fmt"
)

func main() {
	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Berryllum"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"

	fmt.Println(elements["Li"])
	/*
	for i := 0; i > len(elements[]); i++ {
		fmt.Println(elements[i])
	}
	*/
	fmt.Println(elements["Un"])
	name, ok := elements["Un"]
	fmt.Println(name, ok)
}
