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

func (mkt *Stock_Market) ProcessTransaction(ask bool, bid bool, nodeC NodeCollection, ticket *ticket, buySell bool) { // buy is true, sell is false
	// using some weighted (midpoint) average formula to determine
	// stock price
	if buySell { //buying
		if len(mkt.ask.log) == 0 {
			return
		}
		var attemptedPrice Price = 65535
		for ticket.quantity != 0 {
			_, truth := mkt.ask.getLowestPrice()
			if !truth {
				return
			}
			otherTicket := mkt.ask.log[mkt.ask.index]
			node_oppose := nodeC.Nodes[otherTicket.address]
			node := nodeC.Nodes[ticket.address]
			attemptedPrice = otherTicket.price
			price := ticket.price
			if attemptedPrice > ticket.price {
				break
			}
			var numberOfShares int = 0
			numberOfShares = min(ticket.quantity, otherTicket.quantity)
			nodeC.updateNodeInvestmentsFromFilledOrder(numberOfShares, price, node_oppose.Address, otherTicket)
			nodeC.updateNodeInvestmentsFromFilledOrder(numberOfShares, price, node.Address, ticket)
			mkt.ask.editTicket(numberOfShares, ticket)
			mkt.ask.editTicket(numberOfShares, otherTicket)
			mkt.ask.lowestPrice = 65535
		}
	} else { //selling
		if len(mkt.bid.log) == 0 {
			return
		}
		var attemptedPrice Price = 0
		for ticket.quantity == 0 {
			_, truth := mkt.bid.getHighestPrice()
			if !truth {
				return
			}
			otherTicket := mkt.bid.log[mkt.bid.index]
			node_oppose := nodeC.Nodes[otherTicket.address]
			node := nodeC.Nodes[ticket.address]
			attemptedPrice = otherTicket.price
			price := ticket.price
			if attemptedPrice < ticket.price {
				break
			}
			var numberOfShares int = 0
			numberOfShares = min(ticket.quantity, otherTicket.quantity)
			nodeC.updateNodeInvestmentsFromFilledOrder(numberOfShares, price, node_oppose.Address, otherTicket)
			nodeC.updateNodeInvestmentsFromFilledOrder(numberOfShares, price, node.Address, ticket)
			mkt.bid.editTicket(numberOfShares, ticket)
			mkt.bid.editTicket(numberOfShares, otherTicket)
			mkt.bid.highestPrice = 0
		}
	}

	price := (mkt.ask.lowestPrice + mkt.bid.highestPrice) / 2
	if ask && bid {
		mkt.PrevPrices = append(mkt.PrevPrices, &price)
	} else {
		mkt.PrevPrices = append(mkt.PrevPrices, mkt.PrevPrices[len(mkt.PrevPrices)-1])
	}
}

func (mkt *Stock_Market) Buy(ticket *ticket) {
	mkt.bid.log = append(mkt.bid.log, ticket)
	// println("Node", ticket.address, "attempted to buy a trade")
}

func (mkt *Stock_Market) Sell(ticket *ticket) {
	mkt.ask.log = append(mkt.ask.log, ticket)
	// println("Node", ticket.address, "attemped to sell a trade")
}

func (mkt *Stock_Market) OrderToFill(nodeC NodeCollection, ticket *ticket) {
	if ticket.address == 0 {
		return
	}
	lowestAsk, truthBid := mkt.ask.getLowestPrice()
	highestBid, truthAsk := mkt.bid.getHighestPrice()
	if ticket.action { // its buying
		if ticket.price >= lowestAsk && truthBid {
			mkt.ProcessTransaction(truthAsk, truthBid, nodeC, ticket, true)
			// println("order filled")
		}
	} else { // its selling
		if ticket.price <= highestBid && truthAsk {
			mkt.ProcessTransaction(truthAsk, truthBid, nodeC, ticket, false)
			// println("order filled")
		}
	}
}
