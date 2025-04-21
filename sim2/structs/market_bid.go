package structs

type Bid struct {
	log          map[address]*ticket
	highestPrice Price
	quantity     int
	index        address
}

func (bid *Bid) init() {
	bid.highestPrice = 0
	bid.log = map[address]*ticket{}
}

func (bid *Bid) getHighestPrice() (Price, bool) { // Price, is it valid?, quantity, index for the bid
	bid.highestPrice = -1
	bid.index = 0

	if len(bid.log) == 0 {
		return 0, false
	}

	for i, val := range bid.log {
		if val.price > bid.highestPrice {
			bid.highestPrice = bid.log[i].price
			bid.quantity = bid.log[i].quantity
			bid.index = i
		}
	}
	return bid.highestPrice, true
}

func (bid *Bid) editTicket(shares int, tickAdd address) int { // buy is true, sell is false
	ticket := bid.log[tickAdd]
	if ticket.quantity == shares || ticket.quantity == 0 {
		ticket.quantity = 0
		delete(bid.log, tickAdd)
		return ticket.quantity
	} else {
		ticket.quantity -= shares
		return ticket.quantity
	}
}

// func (bid *Bid) removeFromList(tickAdd address) {
// 	delete(bid.log, tickAdd)
// }
