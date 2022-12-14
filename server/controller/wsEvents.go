package controller

import "encoding/json"

type Event struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

const (
	// EventSendMessage is the event name for new chat messages sent
	EventSendMessage = "send_message"
)

type EventHandler func(event Event, client *Client) error
