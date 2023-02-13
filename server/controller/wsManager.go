package controller

import (
	"errors"
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var (
	websocketUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

type Manager struct {
	clients ClientList
	sync.RWMutex
	handlers map[string]EventHandler
}

func SetManager() *Manager {
	m := &Manager{
		clients:  make(ClientList),
		handlers: make(map[string]EventHandler),
	}
	m.SetEventHandler()
	return m
}

func (manager *Manager) SetEventHandler() {
	manager.handlers[EventSendMessage] = func(event Event, client *Client) error {
		fmt.Println(event)
		return nil
	}
}

func (manager *Manager) ServeWS(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received new Connection")
	websocketUpgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	connection, err := websocketUpgrader.Upgrade(w, r, nil)

	if err != nil {
		fmt.Println(err)
		return
	}

	client := setClient(connection, manager)

	manager.AddNewClient(client)
	go client.ReadMessages()
	go client.WriteMessage()

}

func (manager *Manager) AddNewClient(client *Client) {
	manager.Lock()
	defer manager.Unlock()
	manager.clients[client] = true

}

func (manager *Manager) DeleteClient(client *Client) {
	manager.Lock()
	defer manager.Unlock()
	fmt.Println("Removed Client")
	client.Connection.Close()
	delete(manager.clients, client)
}

func (manager *Manager) RouteEvent(event Event, client *Client) error {
	handler, found := manager.handlers[event.Name]

	if found {
		err := handler(event, client)

		if err != nil {
			fmt.Println(err)
		}

		return nil
	} else {
		return errors.New("this event type is not supported")
	}

}
