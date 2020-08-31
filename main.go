package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/sandblox-official/game/server"
)

func main() {

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
	//Set Port
	port := "8080"
	log.Println("Server started at", port)

	//Serve Static Files
	http.Handle("/", http.FileServer(http.Dir("webroot")))
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
