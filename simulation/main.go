package main

import (
	"fmt"
	str "main/structs"
	wbs "main/websockets"
	"net/http"
)

func initWebsocket() {
	http.HandleFunc("/ws", wbs.WsHandler)
	fmt.Println("WebSocket server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func runSim() [][3]float32 {
	var sim str.Simulation
	sim.GENERATIONS = 10
	sim.NEIGHBORS_RADIUS = 2
	sim.TIMESTEP = 0.1
	sim.SEEDNUM = 0
	sim.NUMBER_OF_NODES = 100

	sim.GenerateNodelist()
	sim.UpdateMovement()
	return sim.ReturnLocations()
}

func main() {
	runSim()
}
