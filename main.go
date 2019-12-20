package main

import "fmt"

func main() {
	// get new deck and save it to file
	sDeck := newDeckStruct()

	// test shuffle func
	sDeck.shuffle()
	fmt.Println(sDeck)
}
