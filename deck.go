package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

type Card struct {
	suit  string
	value string
}

type tDeck struct {
	cards []Card
}

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[0:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d tDeck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d.cards {
		newPosition := r.Intn(len(d.cards) - 1)

		d.cards[i], d.cards[newPosition] = d.cards[newPosition], d.cards[i]
	}
}

// changes deck type to a struct
// func deckToStruct(d deck) tDeck {
// 	strDeck := []Card{}
// 	dt := tDeck{strDeck}

// 	for k, v := range d {
// 		item := Card{k, v}
// 		dt.cards = append(dt.cards, item)
// 	}
// 	return dt
// }

// creates new struct of deck, need to depriciate the string method
func newDeckStruct() tDeck {
	cards := []Card{}
	dt := tDeck{cards}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, card := range cardValues {
			dt.cards = append(dt.cards, Card{suit, card})
		}
	}
	return dt
}
