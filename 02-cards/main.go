package main

func main() {
	// cards := newDec()
	// // fmt.Println(cards.toString())
	// cards.saveToFile("my_cards")

	cards := newDeckFromFile("my_cards")
	cards.print()
}
