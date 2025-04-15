package structs

type Stock_Market struct {
	Price      Price
	PrevPrices []*Price
	bid        *Bid
	ask        *Ask
}

func (mkt *Stock_Market) InitializeMarket(init_price Price, bid *Bid, ask *Ask) {
	mkt.Price = init_price
	mkt.PrevPrices = append(mkt.PrevPrices, &init_price)
	mkt.ask = ask
	mkt.bid = bid
	mkt.bid.init()
	mkt.ask.init()
}

func (mkt *Stock_Market) ProcessTransaction() {
	// using some weighted (midpoint) average formula to determine
	// stock price
	var alpha Price = 0.1
	highestBid, truthBid := mkt.bid.getHighestPrice()
	lowestAsk, truthAsk := mkt.ask.getLowestPrice()
	midPrice := (highestBid + lowestAsk) / 2
	// println("part 1", *mkt.PrevPrices[len(mkt.PrevPrices)-1])
	// println("part 2", (1-alpha)*midPrice)
	price := *mkt.PrevPrices[len(mkt.PrevPrices)-1] + (1-alpha)*midPrice
	mkt.Price = price
	// println("yo", midPrice)
	if truthAsk && truthBid {
		mkt.PrevPrices = append(mkt.PrevPrices, &price)
	} else {
		mkt.PrevPrices = append(mkt.PrevPrices, mkt.PrevPrices[len(mkt.PrevPrices)-1])
	}
}

func (mkt *Stock_Market) Buy(ticket ticket) {
	mkt.bid.log = append(mkt.bid.log, &ticket)
	// println("Node", ticket.address, "attempted to buy a trade")
}

func (mkt *Stock_Market) Sell(ticket ticket) {
	mkt.ask.log = append(mkt.ask.log, &ticket)
	// println("Node", ticket.address, "attemped to sell a trade")
}

func (mkt *Stock_Market) OrderToFill(nodeC NodeCollection, ticket ticket) {
	if ticket.address == 0 {
		return
	}
	lowestPrice, truthBid := mkt.ask.getLowestPrice()
	highestPrice, truthAsk := mkt.bid.getHighestPrice()
	if ticket.action { // its buying
		if ticket.price >= lowestPrice && truthBid {
			mkt.ProcessTransaction()
			nodeC.updateNodeInvestmentsFromFilledOrder(ticket)
			mkt.bid.removeFromList(ticket)
			// println("order filled")
		} else {
			mkt.ProcessTransaction()
		}
	} else { // its selling
		if ticket.price <= highestPrice && truthAsk {
			mkt.ProcessTransaction()
			nodeC.updateNodeInvestmentsFromFilledOrder(ticket)
			mkt.ask.removeFromList(ticket)
			// println("order filled")
		} else {
			mkt.ProcessTransaction()
		}
	}
}
