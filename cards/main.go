package main

import "fmt"

func main() {
	cards := newDeck()
	err := cards.saveToFile("deck")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println("Deck saved")
	newDeck := newDeckFromFile("deck")
	fmt.Println(newDeck)
}
