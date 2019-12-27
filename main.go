package main

import "fmt"

func main() {
	// get new deck and save it to file
	sDeck := newDeckStruct()

	// test shuffle func
	sDeck.shuffle(5)
	sDeck.print()

	x := dealX(sDeck, 5, 3)
	for _, v := range x {
		v.print()
		fmt.Println(v.handTotal())
	}

	h1 := hBlkJck(x)
	for _, v := range h1 {
		fmt.Println(v)
	}

	w := compareToDealer(h1[0], h1)
	fmt.Println(w)
}
