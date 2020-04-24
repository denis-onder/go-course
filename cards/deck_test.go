package main

import "testing"

var cards deck

func TestNewDeck(t *testing.T) {
	cards = newDeck()
	if len(cards) != 52 {
		t.Errorf("Expected deck length of 52, but got %d", len(cards))
	}
	if cards[0] != "Ace of Spades" {
		t.Errorf("Expeceted first card to be the Ace of Spades, but got %s", cards[0])
	}
	if cards[len(cards)-1] != "King of Hearts" {
		t.Errorf("Expeceted last card to be the King of Hearts, but got %s", cards[len(cards)-1])
	}
}
