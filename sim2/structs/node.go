package structs

import (
	nodeTypes "main/structs/models"
)

type NodeFunctions interface {
	Propagate()
	UpdateBeliefs()
}

type models struct {
	node_LSTM *nodeTypes.LSTM
	node_MLP  *nodeTypes.MLP
	node_RNN  *nodeTypes.RNN
}

type Node struct {
	trader models
	info   Info
	layers int
}
