package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	s := 52

	if len(d) != s {
		t.Errorf("Expected %v deck length, got %v", s, len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expect first card of ace of spades, got %v", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected last card of King of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")
	s := 52

	if len(loadedDeck) != s {
		t.Errorf("Expected %v cards, got only %v", s, len(loadedDeck))
	}

	os.Remove("_decktesting")
}
