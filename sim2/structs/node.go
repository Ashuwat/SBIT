package structs

import (
	models "main/structs/models"
	"math/rand/v2"
)

type address = int32

type Node struct {
	trader     models.Models
	info       Info
	layers     int
	address    address
	investment price
}

func InitializeNode() *Node {
	n := new(Node)
	n.layers = rand.IntN(5)
	n.address = rand.Int32()
	return n
}

func (node *Node) UpdateInfo(info Info) {
	node.info = info
}

func (node *Node) DecideToTrade(i int, mkt Stock_Market) {
	// true means buy, false meanse sell
	node.trader.MLP.PropagateForward(node.info.information)
	action, value := node.trader.MLP.PropagateBackward()
	var ticket ticket
	ticket.action = ticket.action
	ticket.price = value
	ticket.address = node.address
	ticket.date = i
	if action {
		mkt.Buy(ticket)
	} else {
		mkt.Sell(ticket)
	}
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
