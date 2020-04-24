package main

import "fmt"

func main() {
	cards := []string{newCard(), newCard()}
	cards = append(cards, newCard())
	for i := 0; i < len(cards); i++ {
		fmt.Println(cards[i])
	}
}

func newCard() string {
	return "Five of Diamonds"
}
