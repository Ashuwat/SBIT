package structs

type ticket struct {
	address  address
	date     int
	action   bool
	price    Price
	quantity int
}

func (ticket *ticket) updateTicket() {
	//
}
