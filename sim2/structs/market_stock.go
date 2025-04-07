package structs

import "github.com/x448/float16"

type Price = float16.Float16

type Stock_Market struct {
	Price      Price
	PrevPrices []Price
	bid        bid
	ask        ask
}

func (mkt *Stock_Market) ProcessTransaction() {
	// using some weighted (midpoint) average formula to determing
	// stock price
	var alpha Price = float16.Fromfloat32(0.9)
	highestBid := mkt.bid.getHighestPrice()
	lowestAsk := mkt.ask.getLowestPrice()
	midPrice := (highestBid + lowestAsk) / 2
	println("e", len(mkt.bid.log), len(mkt.ask.log))
	mkt.Price = mkt.PrevPrices[len(mkt.PrevPrices)-1] + (1-alpha)*midPrice
	mkt.PrevPrices = append(mkt.PrevPrices, mkt.Price)
}

func (mkt *Stock_Market) InitializeMarket(init_price Price) {
	mkt.Price = init_price
	mkt.PrevPrices = append(mkt.PrevPrices, init_price)
	mkt.ask.init()
	mkt.bid.init()
}

func (mkt *Stock_Market) Buy(ticket ticket) {
	mkt.bid.log = append(mkt.bid.log, ticket)
	println("Node ", ticket.address, "attempted to buy a trade")
}

func (mkt *Stock_Market) Sell(ticket ticket) {
	mkt.ask.log = append(mkt.ask.log, ticket)
	println("Node ", ticket.address, "attemped to sell a trade")
}

func (mkt *Stock_Market) OrderToFill(nodeC NodeCollection, ticket ticket) {
	if ticket.address == 0 {
		return
	}
	if ticket.action { // its buying
		if ticket.price >= mkt.ask.getLowestPrice() {
			mkt.ProcessTransaction()
			nodeC.updateNodeInvestmentsFromFilledOrder(ticket)
			mkt.ask.removeFromList(ticket)
			println("order filled")
		} else {
			mkt.ProcessTransaction()
		}
	} else { // its selling
		if ticket.price <= mkt.bid.getHighestPrice() {
			mkt.ProcessTransaction()
			nodeC.updateNodeInvestmentsFromFilledOrder(ticket)
			mkt.bid.removeFromList(ticket)
			println("order filled")
		} else {
			mkt.ProcessTransaction()
		}
	}
}
