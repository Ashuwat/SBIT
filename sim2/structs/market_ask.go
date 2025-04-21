package structs

type Ask struct {
	log         map[address]*ticket
	lowestPrice Price
	quantity    int
	index       address
}

func (ask *Ask) init() {
	ask.lowestPrice = 65535
	ask.log = map[address]*ticket{}
}

func (ask *Ask) getLowestPrice() (Price, bool) {
	ask.lowestPrice = 21000
	ask.index = 0

	if len(ask.log) == 0 {
		return 0, false
	}
	for i, val := range ask.log {
		if val.price < ask.lowestPrice {
			ask.lowestPrice = ask.log[i].price
			ask.quantity = ask.log[i].quantity
			ask.index = i
		}
	}
	return ask.lowestPrice, true
}

func (ask *Ask) editTicket(shares int, tickAdd address) int { // buy is true, sell is false
	ticket := ask.log[tickAdd]
	if ticket.quantity == shares || ticket.quantity == 0 {
		ticket.quantity = 0
		delete(ask.log, tickAdd)
		return ticket.quantity
	} else {
		ticket.quantity -= shares
		return ticket.quantity
	}

}

// func (ask *Ask) removeFromList(tickAdd address) {
// 	delete(ask.log, tickAdd)
// }
