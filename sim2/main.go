package main

import (
	"main/json"
	"main/structs"
)

func main() {
	NODE_AMOUNT := 10000
	SIMULATION_EPOCHS := 50
	var INIT_PRICE float32 = 1000
	var output json.Output
	var stock_market structs.Stock_Market
	var node_collection structs.NodeCollection
	var bid structs.Bid
	var ask structs.Ask

	output.Init()
	node_collection.Init()

	stock_market.InitializeMarket(INIT_PRICE, &bid, &ask)
	for i := range NODE_AMOUNT {
		var something = structs.InitializeNode(int64(i), 20, 100)
		node_collection.Nodes[something.Address] = something
	}

	var count int64 = 0
	for i := range SIMULATION_EPOCHS {
		var innerCount int64 = 0
		var info structs.Info
		for address := range node_collection.Nodes {
			node := node_collection.Nodes[address]
			node.UpdateInfo(info)
			// println(1)
			ticket := node.DecideToTrade(count, stock_market)
			// println(2)
			stock_market.OrderToFill(node_collection, ticket)
			// println(3)
			output.AppendNodeStats(innerCount, node)
			// println()
			count++
			innerCount++
		}
		print(i)
	}

	// for i := range node_collection.Nodes {
	// 	println(node_collection.Nodes[i].Investment)
	// }
	println("done!")
	output.AppendAllNodeInvestments(node_collection)
	output.AppendMarketPrices(stock_market)
	json.ConvertJson(output)
}
