package main

import (
	"fmt"
	"os"
	"testing"
)

//*testing.T =>
func TestNewDeck(t *testing.T) {
	fmt.Println("Running TestNewDeck...")
	d := newDec()
	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected last card of King of Clubs, but got %v", d[len(d)-1])
	}

}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	fmt.Println("Running TestSaveToDeckAndNewDeckFromFile...")
	os.Remove("_decktesting")

	deck := newDec()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
