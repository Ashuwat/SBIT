package structs

import (
	"main/structs/models"
	"math/rand/v2"
)

type address = int32

type Node struct {
	trader     models.Models
	info       Info
	layers     int
	address    address
	investment Price
}

func InitializeNode() *Node {
	var beliefs []float32
	var neurons [][2]float32

	n := new(Node)
	n.layers = rand.IntN(5)
	n.address = rand.Int32()
	for range 10 {
		beliefs = append(beliefs, float32(rand.IntN(10)))
		neurons = append(neurons, [2]float32{float32(rand.IntN(10)), float32(rand.IntN(10))})
	}
	n.trader.MLP.InitializeNetwork(beliefs, neurons)
	return n
}

func (node *Node) UpdateInfo(info Info) {
	node.info = info
}

func (node *Node) DecideToTrade(i int, mkt Stock_Market) ticket {
	// true means buy, false meanse sell

	action, value := node.trader.MLP.RandomFunc()
	var ticket ticket
	if action == 1 {
		ticket.action = true // buy
	} else if action == 2 {
		ticket.action = false // sell
	} else {
		return ticket
	}
	ticket.price = value
	ticket.address = node.address
	ticket.date = i
	if ticket.action {
		mkt.Buy(ticket)
	} else {
		mkt.Sell(ticket)
	}
	return ticket
}

// for group functions
type NodeCollection struct{ Nodes []Node }

func (nodeC *NodeCollection) updateNodeInvestmentsFromFilledOrder(ticket ticket) {
	for i := range nodeC.Nodes {
		if nodeC.Nodes[i].address == ticket.address {
			if ticket.action {
				nodeC.Nodes[i].investment -= ticket.price
			}
		}
	}
}
