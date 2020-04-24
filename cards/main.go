package main

import "fmt"

func main() {
	cards := []string{newCard(), newCard()}
	cards = append(cards, newCard())
	// for i := 0; i < len(cards); i++ {
	// 	fmt.Println(cards[i])
	// }
	for i, card := range cards {
		fmt.Printf("%d. %s\n", i, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}
