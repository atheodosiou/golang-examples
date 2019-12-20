package main

import "fmt"

//Create a new type of 'deck'
//wich is a slice of strings
type deck []string

func (d deck) print() { // <= This is a reciever!
	for i, card := range d {
		fmt.Println(i, card)
	}
}
