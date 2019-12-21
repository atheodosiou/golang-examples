package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

//Create a new type of 'deck'
//wich is a slice of strings
type deck []string

//Create a new deck of playing cards
func newDec() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() { // <= This is a reciever!
	for i, card := range d {
		fmt.Println(i, card)
	}
}

//This is how we return multiple values fron a function
func deal(d deck, handSise int) (deck, deck) {
	return d[:handSise], d[handSise:]
}

//Converting a deck to string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		//Option #1 - log the error and return a call to newDeck()
		//Option #2 - Log the error and entirely quit the program
		fmt.Println("ERROR ==>", err)
		os.Exit(1)
	}
	//Convert the byte slice (bs) into a string and then
	//we convert this string into a slice of strings via the split method
	s := strings.Split(string(bs), ",")
	//Finaly, we convert the string slice into a deck with a type conversion
	return deck(s)
}
