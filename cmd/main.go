package main

import (
	"github.com/claudiu-persoiu/tequila-worms-go/internal"
	"github.com/claudiu-persoiu/tequila-worms-go/internal/game"
	"github.com/claudiu-persoiu/tequila-worms-go/internal/websockets"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	sendToHub := make(chan internal.Message, 100)
	sendToGame := make(chan internal.Message, 100)

	g := game.NewGame(sendToHub, sendToGame, 80, 55)

	go g.Start()

	h := websockets.NewHub(sendToHub, sendToGame)

	go h.Run()

	app.Static("/", "./public")

	app.Get("/ws", websocket.New(h.NewConnection()))

	app.Listen(":3000")
}
