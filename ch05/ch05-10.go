package main

import (
	"fmt"
)

func main() {
	favorites := map[string]map[string]string{
		"Dad": map[string]string{
			"Game":"Podracer, Battlezone, or Halo 3",
			"Food":"Probably Pizza",
			"Drink":"likely Orange Juice",
		},
		"Rowan": map[string]string{
			"Game":"Fortnite",
			"Food":"Hamburgers & Tacos",
			"Drink":"and I quote, 'I do not know'",
		},
		"Juniper": map[string]string{
			"Game":"Mario",
			"Food":"Strawberries & Bananas",
			"Drink":"Dr. Pepper",
		},
		"Wren": map[string]string{
			"Game":"Boeboes",
			"Food":"Boeboes",
			"Drink":"Boeboes",
		},
	}

	fmt.Println("\n\nWhat is Dad's favorite game, food, and drink?\n")
	if fav, ok := favorites["Dad"]; ok {
		fmt.Println(fav["Game"], fav["Food"], fav["Drink"])
	}
        fmt.Println("\n\nWhat is Rowan's favorite game, food, and drink?\n")
        if fav, ok := favorites["Rowan"]; ok {
                fmt.Println(fav["Game"], fav["Food"], fav["Drink"])
        }
        fmt.Println("\n\nWhat is Juniper's favorite game, food, and drink?\n")
        if fav, ok := favorites["Juniper"]; ok {
                fmt.Println(fav["Game"], fav["Food"], fav["Drink"])
        }
        fmt.Println("\n\nWhat is Wren's favorite game, food, and drink?\n")
        if fav, ok := favorites["Wren"]; ok {
                fmt.Println(fav["Game"], fav["Food"], fav["Drink"])
        }
}
