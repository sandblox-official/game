package server

import (
	"log"

	"github.com/gorilla/websocket"
)



//HandleMessages ...
func HandleMessages(clients map[*websocket.Conn]bool, broadcast chan Message) {
	for {
		// Grab the next message from the broadcast channel
		msg := <-broadcast
		// Send it out to every client that is currently connected
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
