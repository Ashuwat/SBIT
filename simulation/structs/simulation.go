package structs

import (
	"math/rand/v2"
)

type Simulation struct {
	Timestep    float32
	SeedNum     uint64
	nodelist    []Node
	Radius      float32
	Generations uint16
}

func (sim Simulation) GenerateNodelist(length int) []Node {
	var nodelist []Node
	sdf := rand.NewPCG(sim.SeedNum, 0000)
	random := rand.New(sdf)
	for range length {
		var tempNode Node
		var randomList [len(tempNode.Location)]float32
		for i := 0; i < len(tempNode.Location); i++ {
			randomList[i] = random.Float32()
		}
		tempNode.Sim = sim
		// update entire thing here, you're almost done
		tempNode.Location = randomList
		nodelist = append(nodelist, tempNode)
	}
	sim.nodelist = nodelist
	return nodelist
}
