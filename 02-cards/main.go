package main

func main() {
	//VARIABLE DECLARATIONS
	//var card string ="Ace of Spades"
	// card := "Ace of Spades"

	// card := newCard()
	// fmt.Println(card)

	//Slice Declaration
	cards := newDec()

	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()
}
