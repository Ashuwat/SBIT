package structs

type price = float32

type Stock_Market struct {
	price      price
	prevPrices []price
	bid        bid
	ask        ask
}

func (mkt *Stock_Market) ProcessTransaction() {
	// using some weighted (midpoint) average formula to determing
	// stock price

	var alpha price = 0.9
	midPrice := mkt.bid.getHighestPrice()
	mkt.price = mkt.prevPrices[len(mkt.prevPrices)-2] + (1-alpha)*midPrice

}

func (mkt *Stock_Market) Buy(ticket ticket) {
	mkt.bid.log = append(mkt.bid.log, ticket)
	println("Node ", ticket.address, "bought a trade")
}

func (mkt *Stock_Market) Sell(ticket ticket) {
	mkt.ask.log = append(mkt.ask.log, ticket)
	println("Node ", ticket.address, "bought a trade")
}

func (mkt *Stock_Market) OrderToFill(nodeC NodeCollection, ticket ticket) {
	if ticket.action { // its buying
		if ticket.price >= mkt.ask.getLowestPrice() {
			nodeC.updateNodeInvestmentsFromFilledOrder(ticket)
		}
	} else { // its selling
		if ticket.price <= mkt.bid.getHighestPrice() {
			nodeC.updateNodeInvestmentsFromFilledOrder(ticket)
		}
	}
}
