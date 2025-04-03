package models

type MLP struct {
	layers  int
	belief  []int
	nuerons []int
	input   []float32
}

func (mlp *MLP) InitializeNetwork(belief []int, neurons []int) {
	mlp.belief = belief
	mlp.nuerons = neurons

}

func (mlp *MLP) PropagateForward(belief, vector []int) {

}

func (mlp *MLP) PropagateBackward() {

}
