package main

import "fmt"

/*
	Create a new type 'deck'.
	Deck is a slice of strings
*/
type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Printf("%d. %s\n", i, card)
	}
}
