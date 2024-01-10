package game

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/claudiu-persoiu/tequila-worms-go/internal"
)

type game struct {
	sendToHub  chan internal.Message
	sendToGame chan internal.Message
	players    *repo
	size       size
}

type size struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func NewGame(sendToHub, sendToGame chan internal.Message, x int, y int) *game {
	return &game{
		sendToGame: sendToGame,
		sendToHub:  sendToHub,
		players:    NewRepo(),
		size:       size{X: x, Y: y},
	}
}

func (g *game) Start() {
	go g.handleMessages()

	for range time.Tick(time.Millisecond * 500) {
		if g.players.GetSize() == 0 {
			continue
		}

		for _, w := range g.players.GetWorms() {
			w.Step()

			if checkHitTheWall(w.pieces[0], g.size) || checkHitSelf(w) {
				w.Kill()
				continue
			}

			intersectBetweenWorms(w, g.players.GetWorms())
		}

		// filter dead worm
		for uuid, w := range g.players.GetWorms() {
			if w.IsDead() {
				g.players.RemoveWorm(uuid)
				g.broadcastDeadWorm(w)
				g.sendToHub <- internal.Message{
					UUID:   w.uuid,
					Action: "you dead",
					Data:   true,
				}
			}
		}

		g.broadcastWormsData()
	}
}

func (g *game) handleMessages() {
	for {
		message := <-g.sendToGame
		switch message.Action {
		case "connected":
			g.broadcastWormList()
			g.sendMatrixSize(message.UUID)
			fmt.Println("player connected in game", message.UUID)
		case "disconnected":
			fmt.Println("player disconnected", message.UUID)
			w := g.players.RemoveWorm(message.UUID)
			if w != nil {
				g.broadcastDeadWorm(w)
			}
			g.broadcastWormList()
		case "start game":
			g.players.AddWorm(message.UUID, NewWorm(message.UUID, message.Data.(string), g.getRandomPosition()))

			g.broadcastWormList()
		case "new direction":
			g.players.GetWorm(message.UUID).SetDirection(message.Data.(string))
		default:
			fmt.Println("message not handled: ", message)
		}
	}
}

func (g *game) broadcastDeadWorm(w *worm) {
	g.sendToHub <- internal.Message{
		UUID:   "all",
		Action: "dead worm",
		Data:   w.name + randomDeadMessage(),
	}
}

func (g *game) broadcastWormList() {
	g.sendToHub <- internal.Message{
		UUID:   "all",
		Action: "player list",
		Data:   g.players.GetWormsList(),
	}
}

func (g *game) broadcastWormsData() {
	g.sendToHub <- internal.Message{
		UUID:   "all",
		Action: "worm data",
		Data:   g.players.GetWormsDataList(),
	}
}

func (g *game) sendMatrixSize(uuid string) {
	g.sendToHub <- internal.Message{
		UUID:   uuid,
		Action: "matrix size",
		Data:   g.size,
	}
}

func getRandomWithBezel(x int, bezel int) int {
	return rand.Intn(x-(bezel*2)) + bezel
}

func (g *game) getRandomPosition() piece {
	return piece{
		X: getRandomWithBezel(g.size.X, 2),
		Y: getRandomWithBezel(g.size.Y, 2),
	}
}
