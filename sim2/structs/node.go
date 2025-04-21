package structs

import (
	"main/structs/models"
	"math/rand/v2"
)

type address = int64

type Node struct {
	trader     models.Models
	info       Info
	layers     int
	Address    address
	Investment Price
	Shares     int
}

func InitializeNode(i int64, investment Price, shares int) *Node {
	var beliefs []float32
	var neurons [][2]float32
	n := new(Node)
	n.Investment = investment
	n.Shares = shares
	n.layers = rand.IntN(5)
	n.Address = i
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

func (node *Node) DecideToTrade(i int64, mkt Stock_Market) *ticket {
	// true means buy, false meanse sell

	action, value, limit := node.trader.MLP.RandomFunc()
	ticket := new(ticket)
	if action == 1 {
		ticket.action = true // buy
	} else if action == 2 {
		ticket.action = false // sell
	} else {
		ticket.invest = true
	}

	ticket.price = value
	ticket.address = node.Address
	ticket.tickAdd = int64(i)
	ticket.quantity = limit

	if ticket.action && (node.Investment) >= 0 {
		mkt.Buy(ticket)
		return ticket
	} else if !ticket.action && (node.Shares-ticket.quantity >= 0) {
		mkt.Sell(ticket)
		return ticket
	}
	return nil
}

// for group functions
type NodeCollection struct{ Nodes map[address]*Node }

func (nodeC *NodeCollection) Init() {
	nodeC.Nodes = map[address]*Node{}
}

func (nodeC *NodeCollection) updateNodeInvestmentsFromFilledOrder(shares int, price Price, nodeAdd address, ticket *ticket) {
	if ticket.action {
		nodeC.Nodes[nodeAdd].Investment -= price
		nodeC.Nodes[nodeAdd].Shares += shares
	} else {
		nodeC.Nodes[nodeAdd].Investment += price
		nodeC.Nodes[nodeAdd].Shares -= shares
	}
}
