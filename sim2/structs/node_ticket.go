package structs

type ticket struct {
	address address
	date    int
	action  bool
	price   price
}

func (ticket *ticket) updateTicket() {
	
}
