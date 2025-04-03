package structs

type ask struct {
	log []ticket
}

func (ask *ask) getLowestPrice() price {
	var lowestPrice float32
	for i := range ask.log {
		if ask.log[i].price > lowestPrice {
			lowestPrice = ask.log[i].price
		}
	}
	return lowestPrice
}

func (ask *ask) removeFromList(ticket ticket) {
	for i := range ask.log {
		if ticket == ask.log[i] {
			ask.log = append(ask.log[:i], ask.log[i+1:]...)
		}
	}
}
