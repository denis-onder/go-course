package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

/*
	Create a new type 'deck'.
	Deck is a slice of strings
*/
type deck []string

func newDeck() deck {
	cards := deck{}
	suites := []string{"Spades", "Clubs", "Diamonds", "Hearts"}
	values := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	for _, suite := range suites {
		for _, value := range values {
			cards = append(cards, value+" of "+suite)
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
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	s := string(bytes)
	return deck(strings.Split(s, ","))
}

func (d deck) shuffle() {
	// Create a truly random RNG
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	// Swap indices
	for i := range d {
		newPos := r.Intn(len(d) - 1)
		d[i], d[newPos] = d[newPos], d[i] // Swap syntax
	}
}
