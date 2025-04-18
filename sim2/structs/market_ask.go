package structs

import "slices"

type Ask struct {
	log         []*ticket
	lowestPrice Price
	quantity    int
	index       int
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
			ask.quantity = ask.log[i].quantity
			ask.index = i
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

func (ask *Ask) editTicket(shares int, ticket *ticket) { // buy is true, sell is false
	for i := range ask.log {
		if ticket == ask.log[i] {
			if ticket.quantity == shares {
				ask.log = slices.Delete(ask.log, i, i+1)
				return
			} else {
				ticket.quantity -= shares
				return
			}
		}
	}
}
