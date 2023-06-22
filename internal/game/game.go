package game

import (
	"fmt"

	"github.com/claudiu-persoiu/tequila-worms-go/internal"
	"github.com/google/uuid"
)

type game struct {
	sendToHub  chan internal.Message
	sendToGame chan internal.Message
	players    repo
	size       size
}

type size struct {
	x int
	y int
}

func NewGame(sendToHub, sendToGame chan internal.Message, x int, y int) *game {
	return &game{
		sendToGame: sendToGame,
		sendToHub:  sendToHub,
		size:       size{x: x, y: y},
	}
}

func (g *game) Start() {
	for {
		message := <-g.sendToGame
		switch message.Action {
		case "connected":
			fmt.Println("player connected in game", message.UUID)
		case "disconnected":
			fmt.Println("player disconnected", message.UUID)
			g.players.RemoveWorm(message.UUID)
			//w := g.players.GetWorm(message.UUID)
			//g.sendToHub <- internal.Message{
			//	UUID:   "all",
			//	Action: "dead worm",
			//	Data:   w.name + randomDeadMessage(),
			//}
		case "start game":

		default:
			fmt.Println("message not handled: ", message)
		}
	}
}

func sendMessage(uuid uuid.UUID) {

}
