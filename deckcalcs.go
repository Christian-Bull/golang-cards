package main

type blkJckHand struct {
	cards tDeck
	total int
}

func (h tDeck) handTotal() int {
	var sum int

	for _, v := range h.cards {
		sum += v.value
	}
	return sum
}

func hBlkJck(d []tDeck) []blkJckHand {
	var hands []blkJckHand

	for i := 0; i < len(d); i++ {
		hands = append(hands, blkJckHand{d[i], d[i].handTotal()})
	}
	return hands
}

func compareToDealer(d blkJckHand, p []blkJckHand) []int {
	var win []int

	for _, v := range p {
		if v.total > d.total {
			win = append(win, 2)
		} else if v.total == d.total {
			win = append(win, 1)
		} else {
			win = append(win, 0)
		}
	}
	return win
}
