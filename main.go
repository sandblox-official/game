package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/sandblox-official/game/server"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	//Static Files
	fs := http.FileServer(http.Dir("./webroot"))
	http.Handle("/", fs)
	//Socket Hanlder
	worlds := server.Worlds
	worlds["test1"] = server.CreateWorld()
	go worlds["test1"].Run()
	http.HandleFunc("/test1", func(w http.ResponseWriter, r *http.Request) {
		serveWs(worlds["test1"], w, r)
	})
	worlds["test2"] = server.CreateWorld()
	go worlds["test2"].Run()
	http.HandleFunc("/test2", func(w http.ResponseWriter, r *http.Request) {
		serveWs(worlds["test2"], w, r)
	})
	//Serve and Run Worlds

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func serveWs(world *server.World, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &server.Client{World: world, Conn: conn, Send: make(chan []byte, 256)}
	client.World.Join <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.Emit()
	go client.Consume()

}
