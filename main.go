package main

func main() {
	deck := newDeck()
	deck.saveToFile("./assets/my_deck")

	cards := newDeckFromFile("./assets/my_deck")

	cards.shuffle()
	cards.print()
}
