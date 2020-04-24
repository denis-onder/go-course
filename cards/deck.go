package main

import "fmt"

/*
	Create a new type 'deck'.
	Deck is a slice of strings
*/
type deck []string

func newDeck() deck {
	cards := deck{}
	suites := []string{"Spades", "Clubs", "Diamonds", "Hearts"}
	values := []string{"Ace", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	for _, s := range suites {
		for _, v := range values {
			cards = append(cards, v+" of "+s)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Printf("%d. %s\n", i+1, card)
	}
}

func deal(d deck, size int) (deck, deck) {
	return d[:size], d[size:]
}
