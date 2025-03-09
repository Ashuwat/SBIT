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

func main() {
	var sim str.Simulation
	sim.Generations = 100
	sim.Radius = 2
	sim.Timestep = 0.1
	sim.SeedNum = 4

	// var CurrNodeList = sim.GenerateNodelist(100)
	// var prevNodeList = CurrNodeList
	for i := 0; i < int(sim.Generations); i++ {

	}
	// 	nodelist[0].Node_random_velocity()
	// 	fmt.Println(nodelist[0])
}
