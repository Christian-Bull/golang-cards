package main

import "fmt"

func main() {
	// get new deck and save it to file
	deck := newDeck()
	deck.saveToFile("./assets/my_deck")

	// get the saved file
	cards := newDeckFromFile("./assets/my_deck")

	cards.shuffle()
	cards.print()

	hand, _ := deal(cards, 5)

	fmt.Println(hand)
	sDeck := newDeckStruct()
	fmt.Println(sDeck)

}
