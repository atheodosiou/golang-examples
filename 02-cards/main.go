package main

import "fmt"

func main() {
	//VARIABLE DECLARATIONS
	//var card string ="Ace of Spades"
	// card := "Ace of Spades"

	// card := newCard()
	// fmt.Println(card)

	//Slice Declaration
	cards := []string{"Ace of Spades", newCard()}
	//Append new item on a slice
	cards = append(cards, "Six of Spades")

	//Loops
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}
