package websockets

import (
	"fmt"
	"log"

	"github.com/claudiu-persoiu/tequila-worms-go/internal"
)

type client struct {
	uuid string
	send chan string
}

type messageToHub struct {
	UUID string
	Data string
}

var clients = make(map[string]*client) // Note: although large maps with pointer-like types (e.g. strings) as keys are slow, using pointers themselves as keys is acceptable and fast
var register = make(chan *clientRegisterMessage, 100)
var broadcast = make(chan string, 100)
var unregister = make(chan string, 100)

type hub struct {
	sendToHub  chan internal.Message
	sendToGame chan internal.Message
}

func NewHub(sendToHub, sendToGame chan internal.Message) *hub {
	return &hub{
		sendToGame: sendToGame,
		sendToHub:  sendToHub,
	}
}

func (h *hub) Run() {

	go func() {
		for {
			message := <-h.sendToHub
			m, _ := buildMessage(message.Action, message.Data)
			if message.UUID == "all" {
				broadcast <- m
			} else {
				clients[message.UUID].send <- m
			}
		}
	}()

	for {
		select {
		case c := <-register:
			clients[c.UUID] = &client{
				uuid: c.UUID,
				send: c.Send,
			}
			log.Println("connection registered")
			h.sendToGame <- internal.Message{
				UUID:   c.UUID,
				Action: "connected",
				Data:   "",
			}

		case message := <-broadcast:
			log.Println("message received:", message)
			// Send the message to all clients
			for _, c := range clients {
				c.send <- message
			}

		case uuid := <-unregister:
			// Remove the client from the hub
			delete(clients, uuid)
			// signal disconnect
			h.sendToGame <- internal.Message{
				UUID:   uuid,
				Action: "disconnected",
				Data:   "",
			}

			log.Println("connection unregistered")
		}
	}
}

type playerData struct {
	Name  string `json:"name"`
	Color string `json:"color"`
}

func broadcastPlayerList() {
	res := []playerData{}
	//for _, c := range clients {
	//	if c.Name != "" {
	//		res = append(res, playerData{Name: c.Name, Color: c.Color})
	//	}
	//}

	fmt.Println(res)

	message, _ := buildMessage("player list", res)
	fmt.Println(message)
	broadcast <- message
}
