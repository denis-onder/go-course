package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

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

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("An error has occured! %s", err)
	}
	s := string(bytes)
	return deck(strings.Split(s, ","))
}
