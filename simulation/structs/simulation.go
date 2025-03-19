package structs

import (
	"fmt"
	"math/rand/v2"
)

type Simulation struct {
	TIMESTEP          float32
	SEEDNUM           uint64
	NEIGHBORS_RADIUS  float32
	PROXIMITY_RADIUS  float32
	GENERATIONS       uint16
	previous_nodelist []Node
	current_nodelist  []Node
	NUMBER_OF_NODES   uint32
}

func (sim *Simulation) GenerateNodelist() []Node {
	var nodelist []Node
	sdf := rand.NewPCG(sim.SEEDNUM, 0000)
	random := rand.New(sdf)
	for i := 0; i < int(sim.NUMBER_OF_NODES); i++ {
		var tempNode Node
		var randomList [len(tempNode.Location)]float32
		var data [len(tempNode.data)]float32
		for i := 0; i < len(tempNode.Location); i++ {
			randomList[i] = random.Float32()
		}
		for i := 0; i < len(tempNode.data); i++ {
			data[i] = random.Float32()
		}

		tempNode.Sim = sim
		tempNode.address = uint32(i)
		tempNode.data = data
		tempNode.Location = randomList
		nodelist = append(nodelist, tempNode)
	}
	sim.current_nodelist = nodelist
	sim.previous_nodelist = nodelist
	if sim.current_nodelist != nil || sim.previous_nodelist != nil {
		fmt.Println("successful init of sim!")
	}
	return nodelist
}

func (sim *Simulation) UpdateMovement() {
	for range sim.GENERATIONS {
		for i := 0; i < len(sim.current_nodelist); i++ {
			sim.current_nodelist[i].allMovement()
		}
		sim.previous_nodelist = sim.current_nodelist
		fmt.Println(sim.current_nodelist[0].Location[0])
	}
}

func (sim *Simulation) ReturnLocations() [][3]float32 {
	var locations [][3]float32
	for i := 0; i < len(sim.current_nodelist); i++ {
		locations = append(locations, sim.current_nodelist[i].Location)
	}
	return locations
}
