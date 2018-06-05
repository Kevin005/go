package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool) // connected clients
var broadcast = make(chan Message)           // broadcast channel

// Configure the upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Define our message object
type Message struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Message  string `json:"message"`
}

func main() {
	// Create a simple file server
	http.Handle("/", http.FileServer(http.Dir("./public")))

	// Configure websocket route
	http.HandleFunc("/ws", handleConnections)//WebSocket请求，会实现把http切换为WebSocket

	// Start listening for incoming chat messages
	go handleMessages()

	// Start the server on localhost port 8000 and log any errors
	log.Println("http server started on :9090")
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

/**
1.每当一个浏览器client发起：9090/ws请求时，都会建立一个ws连接,
然后for里边只要有，任何一个client读到了msg就会给broadcast然后，通过range clients 发送给所有的client
 */
func handleConnections(w http.ResponseWriter, r *http.Request) {
	// Upgrade initial GET request to a websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	// Make sure we close the connection when the function returns
	defer ws.Close()

	// Register our new client
	clients[ws] = true //把新的请求端加入到clients cache中去，以后发送数据会同事发给clients cache中所有的client

	for {
		var msg Message
		// Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&msg)//但连接建立后，每当有client消息时，都会在这里ws中拿到
		if err != nil {//WebSocket已经结束
			log.Printf("error1: %v", err)
			delete(clients, ws)
			break
		}
		// Send the newly received message to the broadcast channel
		broadcast <- msg
	}
}

func handleMessages() {
	for {
		// Grab the next message from the broadcast channel
		msg := <-broadcast
		// Send it out to every client that is currently connected
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {//WebSocket已经结束
				log.Printf("error2: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
