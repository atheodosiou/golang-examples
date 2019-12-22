package main

import "fmt"

func main() {
	//Declaring maps

	//First way

	colors := map[string]string{
		"red":   "#ff000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	printMap(colors)
	//Second way
	// var colors map[string]string

	//Third way
	// colors := make(map[string]string)
	// colors["white"] = "#ffffff"

	// Deleting a key value pair from a map
	// delete(colors, "white")

	// fmt.Println(colors)
}

//Iterate over a map
func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
