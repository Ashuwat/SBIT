package main

import (
	"main/structs"
)

func main() {
	NODE_AMOUNT := 50
	SIMULATION := 5
	var INIT_PRICE float32 = 400

	var stock_market structs.Stock_Market
	var node_collection structs.NodeCollection
	stock_market.Price = INIT_PRICE
	stock_market.PrevPrices = append(stock_market.PrevPrices, stock_market.Price)

	for range NODE_AMOUNT {
		var something = structs.InitializeNode()
		node_collection.Nodes = append(node_collection.Nodes, *something)
	}
	for range SIMULATION {
		var info structs.Info
		for i := range node_collection.Nodes {
			node := node_collection.Nodes[i]
			node.UpdateInfo(info)
			ticket := node.DecideToTrade(i, stock_market)
			stock_market.OrderToFill(node_collection, ticket)
		}
	}

	for i := range stock_market.PrevPrices {
		println(int(stock_market.PrevPrices[i]))
	}
}
