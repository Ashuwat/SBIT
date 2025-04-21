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
	if buySell { // buying
		if len(mkt.ask.log) == 0 {
			mkt.PrevPrices = append(mkt.PrevPrices, mkt.PrevPrices[len(mkt.PrevPrices)-1])
		}
		var attemptedPrice Price = 65535
		quant := ticket.quantity
		for quant > 0 {
			_, truth := mkt.ask.getLowestPrice()
			if !truth {
				break
			}
			otherTicket := mkt.ask.log[mkt.ask.index]
			node_oppose := nodeC.Nodes[otherTicket.address]
			node := nodeC.Nodes[ticket.address]
			attemptedPrice = otherTicket.price
			price := ticket.price

			if attemptedPrice >= ticket.price {
				break
			}
			var numberOfShares int = 0
			numberOfShares = min(ticket.quantity, otherTicket.quantity)
			nodeC.updateNodeInvestmentsFromFilledOrder(numberOfShares, price, node_oppose.Address, otherTicket)
			nodeC.updateNodeInvestmentsFromFilledOrder(numberOfShares, price, node.Address, ticket)
			quant = mkt.bid.editTicket(numberOfShares, ticket.tickAdd)
			mkt.ask.editTicket(numberOfShares, otherTicket.tickAdd)
		}
	} else { // selling
		if len(mkt.bid.log) == 0 {
			mkt.PrevPrices = append(mkt.PrevPrices, mkt.PrevPrices[len(mkt.PrevPrices)-1])
		}
		var attemptedPrice Price = 0
		quant := ticket.quantity
		for quant > 0 {
			_, truth := mkt.bid.getHighestPrice()
			if !truth {
				break
			}
			otherTicket := mkt.bid.log[mkt.bid.index]
			node_oppose := nodeC.Nodes[otherTicket.address]
			node := nodeC.Nodes[ticket.address]
			attemptedPrice = otherTicket.price
			price := ticket.price
			if attemptedPrice <= ticket.price {
				break
			}
			var numberOfShares int = 0
			numberOfShares = min(ticket.quantity, otherTicket.quantity)
			nodeC.updateNodeInvestmentsFromFilledOrder(numberOfShares, price, node_oppose.Address, otherTicket)
			nodeC.updateNodeInvestmentsFromFilledOrder(numberOfShares, price, node.Address, ticket)
			quant = mkt.ask.editTicket(numberOfShares, ticket.tickAdd)
			mkt.bid.editTicket(numberOfShares, otherTicket.tickAdd)
		}
	}

	price := (mkt.ask.lowestPrice + mkt.bid.highestPrice) / 2
	if len(mkt.bid.log) > 0 && len(mkt.ask.log) > 0 {
		mkt.PrevPrices = append(mkt.PrevPrices, &price)
	} else {
		mkt.PrevPrices = append(mkt.PrevPrices, mkt.PrevPrices[len(mkt.PrevPrices)-1])
	}
}

func (mkt *Stock_Market) Buy(ticket *ticket) {
	mkt.bid.log[ticket.tickAdd] = ticket
	// println("Node", ticket.address, "attempted to buy a trade")
}

func (mkt *Stock_Market) Sell(ticket *ticket) {
	mkt.ask.log[ticket.tickAdd] = ticket
	// println("Node", ticket.address, "attemped to sell a trade")
}

func (mkt *Stock_Market) OrderToFill(nodeC NodeCollection, ticket *ticket) {
	if ticket == nil {
		return
	}
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
