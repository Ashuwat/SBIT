package structs

import "slices"

type Bid struct {
	log          []*ticket
	highestPrice Price
}

func (bid *Bid) init() {	
	bid.highestPrice = 0
	bid.log = []*ticket{}
}

func (bid *Bid) getHighestPrice() (Price, bool) {
	if len(bid.log) == 0 {
		return 0, false
	}

	for i := range bid.log {
		if bid.log[i].price > bid.highestPrice {
			bid.highestPrice = bid.log[i].price
		}
	}
	return bid.highestPrice, true
}

func (bid *Bid) removeFromList(ticket ticket) {
	for i := range bid.log {
		if ticket == *bid.log[i] {
			bid.log = slices.Delete(bid.log, i, i+1)
			return
		}
	}
}
