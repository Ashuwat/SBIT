package structs

import "slices"

type Ask struct {
	log         []*ticket
	lowestPrice Price
}

func (ask *Ask) init() {
	ask.lowestPrice = 65535
	ask.log = []*ticket{}
}

func (ask *Ask) getLowestPrice() (Price, bool) {
	if len(ask.log) == 0 {
		return 0, false
	}
	for i := range ask.log {
		if ask.log[i].price < ask.lowestPrice {
			ask.lowestPrice = ask.log[i].price
		}
	}
	return ask.lowestPrice, true
}

func (ask *Ask) removeFromList(ticket ticket) {
	for i := range ask.log {
		if ticket == *ask.log[i] {
			ask.log = slices.Delete(ask.log, i, i+1)
			return
		}
	}
}
