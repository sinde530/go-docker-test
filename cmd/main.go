package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var clients = make([]*websocket.Conn, 0)

func main() {
	r := gin.Default()
	r.GET("/ping", pingTest)
	r.GET("/ws", handleWebSocket)
	addr := "/8080"

	fmt.Printf("Server is listening on %s...\n", addr)
	log.Fatal(http.ListenAndServe(":8080", r))
}

func pingTest(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

func handleWebSocket(c *gin.Context) {
	// Upgrade HTTP connection to WebSocket
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("WebSocket upgrade error:", err)
		return
	}

	clients = append(clients, conn)

	// Handle WebSocket messages
	for {
		// Read message from client
		_, msgBytes, err := conn.ReadMessage()
		if err != nil {
			log.Println("WebSocket read error:", err)
			break
		}

		// Broadcast message to all connected clients
		for _, client := range clients {
			err := client.WriteMessage(websocket.TextMessage, msgBytes)
			if err != nil {
				log.Println("WebSocket write error:", err)
			}
		}
	}

	// Close WebSocket connection
	conn.Close()

	// Remove the connection from the clients list
	for i := range clients {
		if clients[i] == conn {
			clients = append(clients[:i], clients[i+1:]...)
			break
		}
	}
}
