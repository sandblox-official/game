package main

import (
	"encoding/json"
	"log"

	"github.com/sandblox-official/game/server"
)

func main() {
	port := ":8080"
	log.Println("Server started at", port)
	reqString := `{"method":"play"}`
	var req server.Request
	err := json.Unmarshal([]byte(reqString), &req)
	if err != nil {
		log.Fatalln("Json Unmarshal", err)
	}
	resp := server.Evaluate(req)
	log.Println(resp)

}
