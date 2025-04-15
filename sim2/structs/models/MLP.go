package models

import (
	"math"
	"math/rand/v2"
)

type MLP struct {
	layers  int
	Belief  []float32
	Neurons [][2]float32
	input   []float32
}

func sigmoid(number float32) float32 {
	return (1 / (1 + float32(math.Exp(float64(number)*-1))))
}

func (mlp *MLP) InitializeNetwork(belief []float32, neurons [][2]float32) {
	mlp.Belief = belief
	mlp.Neurons = neurons
}

func (mlp *MLP) RandomFunc() (int8, float32) {
	x := rand.Float32()
	y := rand.IntN(10)
	if x < 0.45 {
		return 1, float32(y) // buy
	} else if x < 0.9 {
		return 2, float32(y) // sell
	} else {
		return 0, float32(y)
	}
}

func (mlp *MLP) PropagateForward() (int8, float32) {
	var resultingVector_1 float32
	var resultingVector_2 float32
	for i := range mlp.Belief {
		resultingVector_1 = resultingVector_1 + mlp.Belief[i]*mlp.Neurons[i][0]
		resultingVector_2 = resultingVector_2 + mlp.Belief[i]*mlp.Neurons[i][1]
	}
	first := math.Sin(float64(resultingVector_1))
	if first < -0.2 {
		return 0, resultingVector_2 // nothing
	} else if first < 0.2 {
		return 1, resultingVector_2 // buy
	} else {
		return 2, resultingVector_2 // sell
	}
}

func (mlp *MLP) PropagateBackward() {

}

func (mlp *MLP) CovariateAnalysis() {
	
}
