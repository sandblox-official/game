package main

import (
	"github.com/sandblox-official/game/server"
	"encoding/json"
	"log"
)

func main() {
	myJSONString := `{"method":"play"}`
	var myStoredVariable server.Request
	json.Unmarshal([]byte(myJSONString), &myStoredVariable)
	log.Println(myStoredVariable.Method)
}
