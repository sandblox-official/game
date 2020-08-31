package server

import (
	"log"
	"unicode"

	"github.com/gorilla/websocket"
)

//World ...
type World struct {
	Name    string
	Owner   string
	Chats   []Chat
	Players []Player
}

//Player ...
type Player struct {
	Name     string
	Position Position
	Conn     *websocket.Conn
}

//Position ...
type Position struct {
	x float32
	y int
	z int
}

//Chat ...
type Chat struct {
	author string
	body   string
}

//CreateWorld creates a World instance
func CreateWorld(name, owner string) *World {
	log.Println("World", name, "was created")
	var world = World{}
	world.Name = name
	world.Owner = owner
	return &world
}

//Populate ...
func (world *World) Populate(player Player) error {
	world.Players = append(world.Players, player)
	log.Println(player.Name, "entered", world.Name)
	return nil
}

//Valid makes sure each player's data is sufficient
func (player *Player) Valid() bool {
	for i, r := range player.Name {
		if !unicode.IsLetter(r) || i > 16 || i < 3 {
			return false
		}
	}
	return true
}
