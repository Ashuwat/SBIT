package structs

type bid struct {
	log []ticket
}

func (bid *bid) getHighestPrice() price {
	var highestPrice float32
	for i := range bid.log {
		if bid.log[i].price > highestPrice {
			highestPrice = bid.log[i].price
		}
	}
	return highestPrice
}

func (bid *bid) removeFromList(ticket ticket) {
	for i := range bid.log {
		if ticket == bid.log[i] {
			bid.log = append(bid.log[:i], bid.log[i+1:]...)
		}
	}
}
