package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()
	s := 52

	if len(d) != s {
		t.Errorf("Expected %v deck length, got %v", s, len(d))
	}
}
