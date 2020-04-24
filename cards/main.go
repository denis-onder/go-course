package main

import "fmt"

func main() {
	cards := newDeck()
	err := cards.saveToFile("deck")
	if err != nil {
		fmt.Printf("An error has occured! %s", err)
	}
	fmt.Println("Deck saved")
	newDeck := newDeckFromFile("deck")
	fmt.Println(newDeck)
}
