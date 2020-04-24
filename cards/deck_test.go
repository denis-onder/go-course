package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	cards := newDeck()
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

func TestSavingAndReadingDeckFiles(t *testing.T) {
	cards := newDeck()
	filename := "testing_deck"
	if err := cards.saveToFile(filename); err != nil {
		t.Errorf("An error has occured whilst attempting to save the deck: %s", err)
	}
	newCards := newDeckFromFile(filename)
	if len(newCards) != 52 {
		t.Errorf("Expected deck length of 52, but got %d", len(newCards))
	}
	if err := os.Remove(filename); err != nil {
		t.Errorf("An error has occured while removing the deck: %s", err)
	}
}
