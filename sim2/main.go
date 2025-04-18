package main

import (
	"main/json"
	"main/structs"
)

func main() {
	NODE_AMOUNT := 10000
	SIMULATION_EPOCHS := 100
	var INIT_PRICE float32 = 1000
	var output json.Output
	var stock_market structs.Stock_Market
	var node_collection structs.NodeCollection
	var bid structs.Bid
	var ask structs.Ask

	output.Init()
	node_collection.Init()

	stock_market.InitializeMarket(INIT_PRICE, &bid, &ask)
	for range NODE_AMOUNT {
		var something = structs.InitializeNode(20, 100)
		node_collection.Nodes[something.Address] = something
	}

	for i := range SIMULATION_EPOCHS {
		var count int = 0
		var info structs.Info
		for address := range node_collection.Nodes {
			node := node_collection.Nodes[address]
			node.UpdateInfo(info)
			ticket := node.DecideToTrade(count, stock_market)
			stock_market.OrderToFill(node_collection, ticket)
			output.AppendNodeStats(count, node)
			count++
		}
		if (i % 10) == 0 {
			println(i)
		}
	}

	// for i := range node_collection.Nodes {
	// 	println(node_collection.Nodes[i].Investment)
	// }
	println("done!")
	output.AppendAllNodeInvestments(node_collection)
	output.AppendMarketPrices(stock_market)
	json.ConvertJson(output)
}
