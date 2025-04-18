package structs

import "slices"

type Bid struct {
	log          []*ticket
	highestPrice Price
	quantity     int
	index        int
}

func (bid *Bid) init() {
	bid.highestPrice = 0
	bid.log = []*ticket{}
}

func (bid *Bid) getHighestPrice() (Price, bool) { // Price, is it valid?, quantity, index for the bid
	if len(bid.log) == 0 {
		return 0, false
	}

	for i := range bid.log {
		if bid.log[i].price > bid.highestPrice {
			bid.highestPrice = bid.log[i].price
			bid.quantity = bid.log[i].quantity
			bid.index = i
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

func (bid *Bid) editTicket(shares int, ticket *ticket) { // buy is true, sell is false
	for i := range bid.log {
		if ticket == bid.log[i] {
			if ticket.quantity == shares || ticket.quantity == 0 {
				bid.log = slices.Delete(bid.log, i, i+1)
				return
			} else {
				ticket.quantity -= shares
				return
			}
		}
	}
}
