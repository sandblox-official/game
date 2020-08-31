package server

import "net/http"

//CreateClient gives a client object
//that you can write to
func CreateClient(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Websockets not yet functional"))
}
