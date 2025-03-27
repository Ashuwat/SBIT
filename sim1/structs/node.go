package structs

import (
	"math"
	"math/rand/v2"
)

type Node struct {
	data [5]float32
	// credibility (n x 1) matrix
	Location [3]float32
	Velocity [3]float32
	Sim      *Simulation
	trust    map[uint32]float32
	// status - is deterministic of trust (eq) - further work on
	address uint32
	// release for information - latest piece - represented with time
	// release factor - a number/threshold the release has to reach
}

func (n *Node) node_random_velocity() [3]float32 {
	seed := rand.NewPCG(n.Sim.SEEDNUM, 0000)
	random := rand.New(seed)
	var newVelocity [len(n.Location)]float32
	for i := 0; i < len(n.Location); i++ {
		var tempRand = random.Float32()
		newVelocity[i] = tempRand * n.Sim.TIMESTEP
	}
	return newVelocity
}

func (n *Node) nearest_neighbors(radius float32, passDistance bool) ([]Node, [][3]float32) {
	if radius < 0 {
		panic("radius is a neg number")
	}
	var x = n.Location[0]
	var y = n.Location[1]
	var z = n.Location[2]
	var nlist = n.Sim.previous_nodelist
	var nearest_neighbors []Node
	var distance_list [][3]float32
	for i := 0; i < len(n.Sim.previous_nodelist); i++ {
		if nlist[i].Location[0] > (x+radius) && nlist[i].Location[0] < (x-radius) {
			continue
		}
		if nlist[i].Location[1] > (y+radius) && nlist[i].Location[1] < (y-radius) {
			continue
		}
		var x_dist = math.Pow(float64(x-nlist[i].Location[0]), 2)
		var y_dist = math.Pow(float64(y-nlist[i].Location[1]), 2)
		var z_dist = math.Pow(float64(z-nlist[i].Location[2]), 2)
		var temp_distance float32 = float32(x_dist + y_dist + z_dist)
		if temp_distance < radius*radius {
			nearest_neighbors = append(nearest_neighbors, nlist[i])
		}
		if passDistance {
			distance_list = append(distance_list, [3]float32{float32(x_dist), float32(y_dist), float32(z_dist)})
		}
	}
	return nearest_neighbors, distance_list
}

func (n *Node) trust_based_movement(nearest_neighbors []Node, distanceList [][3]float32) [3]float32 {
	if len(n.trust) == 0 {
		return [3]float32{0, 0, 0}
	}
	var x float32
	var y float32
	var z float32

	for i := 0; i < len(nearest_neighbors); i++ {
		val, check := n.trust[nearest_neighbors[i].address]
		if check {
			tempx := distanceList[0][0]
			tempy := distanceList[0][1]
			tempz := distanceList[0][2]

			x = x + tempx*val
			y = y + tempy*val
			z = z + tempz*val
		}
	}
	x = x * n.Sim.TIMESTEP
	y = y * n.Sim.TIMESTEP
	z = z * n.Sim.TIMESTEP
	return [3]float32{x, y, z}
}

func (n *Node) allMovement() {
	velocity := n.node_random_velocity()
	nearest_neighbors, distance_list := n.nearest_neighbors(n.Sim.NEIGHBORS_RADIUS, true)
	trust_velocity := n.trust_based_movement(nearest_neighbors, distance_list)
	for i := 0; i < len(n.Location); i++ {
		n.Location[i] = n.Location[i] + velocity[i] + trust_velocity[i]
		n.Velocity[i] = velocity[i] + trust_velocity[i]
	}
}

func (n *Node) trustFunction(given_node Node) float32 {
	var sum float32
	for i := 0; i < len(given_node.data); i++ {
		sum = sum + (n.data[i]-given_node.data[i])*(n.data[i]-given_node.data[i])
	}
	var initTrust = 1 - sum/float32(len(given_node.data))
	var trust = float32(math.Cos(1 - float64(initTrust)))
	return trust
}
