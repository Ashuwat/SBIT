package models

import "math"

type MLP struct {
	layers  int
	Belief  []float32
	Neurons [][2]float32
	input   []float32
}

func sigmoid(number float32) float32 {
	return (1 / (1 + float32(math.Exp(-1*float64(number)))))
}

func (mlp *MLP) InitializeNetwork(belief []float32, neurons [][2]float32) {
	mlp.Belief = belief
	mlp.Neurons = neurons
}

func (mlp *MLP) PropagateForward() (int8, float32) {
	print()
	var resultingVector_1 float32
	var resultingVector_2 float32
	for i := range mlp.Belief {
		resultingVector_1 = resultingVector_1 + mlp.Belief[i]*mlp.Neurons[i][0]
		resultingVector_2 = resultingVector_2 + mlp.Belief[i]*mlp.Neurons[i][1]
	}
	first := sigmoid(resultingVector_1)
	if first < -0.2 {
		return 0, resultingVector_2
	} else if first < 0.2 {
		return 1, resultingVector_2
	} else {
		return 2, resultingVector_2
	}
}

func (mlp *MLP) PropagateBackward() {

}
