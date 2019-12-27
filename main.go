package main

import "fmt"

func main() {
	// get new deck and save it to file
	sDeck := newDeckStruct()

	// test shuffle func
	sDeck.shuffle(5)
	sDeck.print()

	x := dealX(sDeck, 5, 3)
	fmt.Println(x)
}
