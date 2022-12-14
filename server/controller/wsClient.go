package controller

import (
	"fmt"

	"github.com/gorilla/websocket"
)

type ClientList map[*Client]bool

type Client struct {
	Connection *websocket.Conn
	Manager    *Manager
	Channel    chan []byte
}

func setClient(connection *websocket.Conn, manager *Manager) *Client {

	return &Client{
		Connection: connection,
		Manager:    manager,
		Channel:    make(chan []byte),
	}

}

func (client *Client) ReadMessages() {
	defer func() {
		client.Manager.DeleteClient(client)
	}()

	for {
		mType, payload, err := client.Connection.ReadMessage()

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseAbnormalClosure, websocket.CloseGoingAway) {
				fmt.Println(err)
			}
			break
		}

		fmt.Printf("MessageType: %v, Payload: %s", mType, payload)

		for wsclient := range client.Manager.clients {
			wsclient.Channel <- payload
		}

	}

}

func (client *Client) WriteMessage() {

	defer func() {
		client.Manager.DeleteClient(client)
	}()

	for {

		message, ok := <-client.Channel

		fmt.Println(message)
		if ok {
			err := client.Connection.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				fmt.Println(err)
			}
		}

	}
}
