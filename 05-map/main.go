package main

import "fmt"

func main() {
	//Declaring maps

	//First way

	// colors := map[string]string{
	// 	"red":   "#ff000",
	// 	"green": "4bf745",
	// }

	//Second way
	// var colors map[string]string

	//Third way
	colors := make(map[string]string)
	colors["white"] = "#ffffff"
	fmt.Println(colors)
}
