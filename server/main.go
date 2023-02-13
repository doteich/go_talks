package main

import (
	"go_talks_server/controller"
	"net/http"
)

func main() {
	RouteHandler()
	http.ListenAndServe("127.0.0.1:3001", nil)

}

func RouteHandler() {
	manager := controller.SetManager()
	http.HandleFunc("/ws", manager.ServeWS)
}
