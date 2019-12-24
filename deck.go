package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Create a new type of 'deck'
// which is a slice of strings

type Card struct {
	suit  string
	card  string
	value int
}

type tDeck struct {
	cards []Card
}

func (d tDeck) print() {
	for i, card := range d.cards {
		fmt.Println(i, card)
	}
}

func deal(d tDeck, handSize int) (tDeck, tDeck) {
	var hand tDeck
	for i := 0; i < handSize; i++ {
		hand.cards = append(hand.cards, d.cards[i])
	}
	d.cards = d.cards[handSize:]
	return hand, d
}

func (d tDeck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d.cards {
		newPosition := r.Intn(len(d.cards) - 1)

		d.cards[i], d.cards[newPosition] = d.cards[newPosition], d.cards[i]
	}
}

// creates new struct of deck
func newDeckStruct() tDeck {
	cards := []Card{}
	dt := tDeck{cards}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardNames := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	cardValues := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10, 10, 11}

	for _, suit := range cardSuits {
		for i := 0; i < len(cardNames); i++ {
			dt.cards = append(dt.cards, Card{suit, cardNames[i], cardValues[i]})
		}
	}
	return dt
}
