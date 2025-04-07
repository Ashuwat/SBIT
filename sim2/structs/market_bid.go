package structs

type bid struct {
	log          []ticket
	highestPrice Price
}

func (bid *bid) init() {
	bid.highestPrice = 0
}

func (bid *bid) getHighestPrice() Price {
	// make it so that it gives the best price for a price.
	for i := range bid.log {
		if bid.log[i].price > bid.highestPrice {
			bid.highestPrice = bid.log[i].price
		}
	}
	return bid.highestPrice
}

func (bid *bid) removeFromList(ticket ticket) {
	for i := range bid.log {
		if ticket == bid.log[i] {
			bid.log = append(bid.log[:i], bid.log[i+1:]...)
			return
		}
	}
}
