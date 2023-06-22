package websockets

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/claudiu-persoiu/tequila-worms-go/internal"
	"github.com/gofiber/contrib/websocket"
	uuid2 "github.com/google/uuid"
)

type clientRegisterMessage struct {
	UUID string
	Send chan string
}

func (h *hub) NewConnection() func(*websocket.Conn) {
	return func(c *websocket.Conn) {
		uuid := uuid2.New().String()

		defer func() {
			unregister <- uuid
			c.Close()
		}()

		sendChan := make(chan string)
		mu := sync.Mutex{}
		isClosing := false

		// Register the client
		register <- &clientRegisterMessage{
			UUID: uuid,
			Send: sendChan,
		}

		go func() {
			for {
				message := <-sendChan

				mu.Lock()
				if isClosing {
					return
				}
				if err := c.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
					isClosing = true
					log.Println("write error:", err)

					c.WriteMessage(websocket.CloseMessage, []byte{})
					c.Close()
					unregister <- uuid
				}
				mu.Unlock()
			}
		}()

		for {
			messageType, message, err := c.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					log.Println("read error:", err)
				}

				return // Calls the deferred function, i.e. closes the connection on error
			}

			if messageType == websocket.TextMessage {
				m, err := processMessage(message)
				if err != nil {
					log.Println("error deconding message", message)
				} else {
					h.sendToGame <- internal.Message{
						UUID:   uuid,
						Action: m.Action,
						Data:   m.Value,
					}
				}
			} else {
				log.Println("websocket message received of type", messageType)
			}
		}
	}

}

type message struct {
	Action string `json:"action"`
	Value  any    `json:"value"`
}

func buildMessage(action string, value any) (string, error) {
	r := message{
		Action: action,
		Value:  value,
	}
	rm, err := json.Marshal(r)
	return string(rm), err
}

func processMessage(str []byte) (m message, err error) {
	err = json.Unmarshal(str, &m)
	return m, err
}
