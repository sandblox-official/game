package server

//Request is the json struct for incoming packets
type Request struct {
	Method      string
	Description string
}

//Evaluate is what decides what to do when a request comes in
func Evaluate(req Request) Response {
	switch req.Method {
	case "play":
		resp := Response{}
		return resp
	}
	return Response{Method: "err"}
}
