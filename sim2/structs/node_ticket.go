package structs

type ticket struct {
	address  address
	tickAdd  address
	action   bool
	price    Price
	quantity int
	invest   bool
}

func (ticket *ticket) updateTicket() {
	//
}
