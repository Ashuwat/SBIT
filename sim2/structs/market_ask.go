package structs

type ask struct {
	log         []ticket
	lowestPrice Price
}

func (ask *ask) init() {
	ask.lowestPrice = 65535
}

func (ask *ask) getLowestPrice() Price {
	for i := 0; i < len(ask.log); i++ {
		if ask.log[i].price > ask.lowestPrice {
			ask.lowestPrice = ask.log[i].price
		}
	}
	return ask.lowestPrice
}

func (ask *ask) removeFromList(ticket ticket) {
	for i := range ask.log {
		if ticket == ask.log[i] {
			ask.log = append(ask.log[:i], ask.log[i+1:]...)
		}
	}
}
