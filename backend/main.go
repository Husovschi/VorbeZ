package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

type Message struct {
	Content string `json:"content"`
}

type MessageWithSender struct {
	Message
	Sender *websocket.Conn `json:"-"`
}

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan MessageWithSender)
var upgrader = websocket.Upgrader{}

func main() {
	port := "8080"
	if len(os.Args) > 1 {
		port = os.Args[1]
	}

	http.HandleFunc("/ws", handleConnections)
	go handleMessages()

	log.Printf("WebSocket server is running on ws://vorbez.husovschi.eu:%s\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	clients[ws] = true
	log.Println("New WebSocket connection")

	for {
		var msg Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("Error: %v", err)
			delete(clients, ws)
			break
		}
		log.Printf("Message received: %s", msg.Content)
		broadcast <- MessageWithSender{Message: msg, Sender: ws}
	}
}

func handleMessages() {
	for {
		msgWithSender := <-broadcast
		for client := range clients {
			// Skip the sender
			if client != msgWithSender.Sender {
				err := client.WriteJSON(msgWithSender.Message)
				if err != nil {
					log.Printf("Error: %v", err)
					client.Close()
					delete(clients, client)
				}
			}
		}
	}
}
