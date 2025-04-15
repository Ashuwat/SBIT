package main

import (
	"main/json"
	"main/structs"
)

func main() {
	NODE_AMOUNT := 2500
	SIMULATION_EPOCHS := 10000
	var INIT_PRICE float32 = 400
	var output json.Output
	var stock_market structs.Stock_Market
	var node_collection structs.NodeCollection
	var bid structs.Bid
	var ask structs.Ask

	stock_market.InitializeMarket(INIT_PRICE, &bid, &ask)
	for range NODE_AMOUNT {
		var something = structs.InitializeNode()
		node_collection.Nodes = append(node_collection.Nodes, *something)
	}
	for range SIMULATION_EPOCHS {
		var info structs.Info
		for i := range node_collection.Nodes {
			node := node_collection.Nodes[i]
			node.UpdateInfo(info)
			ticket := node.DecideToTrade(i, stock_market)
			stock_market.OrderToFill(node_collection, ticket)
		}
	}

	// for i := range node_collection.Nodes {
	// 	println(node_collection.Nodes[i].Investment)
	// }

	output.AppendAllNodeInvestments(node_collection)
	json.ConvertJson(output)
}
