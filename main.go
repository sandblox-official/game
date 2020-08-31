package main

import (
	"encoding/json"
	"log"

	"github.com/sandblox-official/game/server"
)

func main() {
	port := ":8080"
	log.Println("Server started at", port)

	//Process Request
	reqString := `{"method":"play"}`
	var req server.Request
	err := json.Unmarshal([]byte(reqString), &req)
	if err != nil {
		log.Fatalln("Json Unmarshal", err)
	}

	//Create Response
	resp, err := server.Evaluate(req)
	if err != nil {
		log.Println(err)
	}
	respRaw, err := json.Marshal(resp)
	if err != nil {
		log.Println("Convert response to json:", err)
	}
	log.Println("Response->", string(respRaw))

}
