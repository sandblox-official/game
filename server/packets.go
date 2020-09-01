package server

//Packet is the data format for socket broadcasts
type Packet struct {
	Method string `json:"method"`
	Data   Data
}

//Data ...
type Data struct {
	//Outgoing
	Player Player
	Chat   Chat
}

//Player ...
type Player struct {
	Name string
	X    float32
	Y    int
	Z    int
}

//Chat ...
type Chat struct {
	From string
	Body string
}

//GetOutputPacket takes an input to generate an output
func (inPacket *Packet) GetOutputPacket() Packet {
	switch inPacket.Method {
	case "move":
		outPacket := &Packet{}
		outPacket.Method = "move"
		outPacket.Data.Player = inPacket.Data.Player
		outPacket.Data.Player.X = inPacket.Data.Player.X
		outPacket.Data.Player.Y = inPacket.Data.Player.Y
		outPacket.Data.Player.Z = inPacket.Data.Player.Z
		return *outPacket

	}
	errPacket := &Packet{
		Method: "error",
	}
	return *errPacket
}
