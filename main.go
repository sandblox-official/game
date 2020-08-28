package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/sandblox-official/game/client"
	"github.com/sandblox-official/game/server"
)

// broadcast channel

// Define our message object

func main() {
	// Create a simple file server
	fs := http.FileServer(http.Dir("./webroot"))
	http.Handle("/", fs)

	// Configure websocket route
	var clients = make(map[*websocket.Conn]bool) // connected clients
	var broadcast = make(chan server.Message)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) { server.HandleConnections(w, r, clients, broadcast) })

	// Start listening for incoming chat messages
	go client.HandleMessages(clients, broadcast)

	// Start the server on localhost port 8000 and log any errors
	log.Println("http server started on :8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
