package main

import "fmt"

func main() {
	//VARIABLE DECLARATIONS
	//var card string ="Ace of Spades"
	// card := "Ace of Spades"
	card := newCard()
	fmt.Println(card)
}

func newCard() string {
	return "Five of Diamonds"
}
