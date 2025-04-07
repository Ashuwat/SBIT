package structs

type ticket struct {
	address address
	date    int
	action  bool
	price   Price
}

func (ticket *ticket) updateTicket() {
	//
}
