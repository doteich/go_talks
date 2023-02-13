package controller

import "time"

type Event struct {
	Name      string `json:"eventName"`
	Payload   string `json:"payload"`
	User      string `json:"user"`
	Timestamp time.Time
}

const (
	// EventSendMessage is the event name for new chat messages sent
	EventSendMessage = "send_message"
)

type EventHandler func(event Event, client *Client) error
