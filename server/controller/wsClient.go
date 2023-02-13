package controller

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gorilla/websocket"
)

type ClientList map[*Client]bool

type Client struct {
	Connection *websocket.Conn
	Manager    *Manager
	Channel    chan Event
}

func setClient(connection *websocket.Conn, manager *Manager) *Client {

	return &Client{
		Connection: connection,
		Manager:    manager,
		Channel:    make(chan Event),
	}

}

func (client *Client) ReadMessages() {
	defer func() {
		client.Manager.DeleteClient(client)
	}()

	for {
		_, data, err := client.Connection.ReadMessage()

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseAbnormalClosure, websocket.CloseGoingAway) {
				fmt.Println(err)
			}
			break
		}
		var newEvent Event

		err = json.Unmarshal(data, &newEvent)

		if err != nil {
			fmt.Println(err)
		}

		newEvent.Timestamp = time.Now().Local()

		fmt.Printf("Payload: %s", newEvent)

		for wsclient := range client.Manager.clients {
			wsclient.Channel <- newEvent
		}

	}

}

func (client *Client) WriteMessage() {

	defer func() {
		client.Manager.DeleteClient(client)
	}()

	for {

		eventObject, ok := <-client.Channel

		data, err := json.Marshal(eventObject)

		if err != nil {
			fmt.Println(err)
		}

		if ok {
			err := client.Connection.WriteMessage(websocket.TextMessage, []byte(data))
			if err != nil {
				fmt.Println(err)
			}
		} else {
			client.Connection.WriteMessage(websocket.CloseMessage, nil)
		}

	}
}
