package main

import (
	"fmt"
	wbs "main/websockets"
	"net/http"
)

func main() {
	http.HandleFunc("/ws", wbs.WsHandler)
	fmt.Println("WebSocket server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
