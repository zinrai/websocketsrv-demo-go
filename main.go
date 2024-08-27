package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func randomString(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	s := make([]rune, n)
	for i := range s {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			fmt.Printf("Error generating random number: %v\n", err)
			return ""
		}
		s[i] = letters[num.Int64()]
	}
	return string(s)
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("Error upgrading to websocket: %v\n", err)
		return
	}
	defer ws.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, _, err := ws.ReadMessage()
			if err != nil {
				fmt.Printf("Error reading message: %v\n", err)
				return
			}
		}
	}()

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			fmt.Println("WebSocket connection closed.")
			return
		case <-ticker.C:
			message := randomString(10)
			err := ws.WriteMessage(websocket.TextMessage, []byte(message))
			if err != nil {
				fmt.Printf("Error writing message: %v\n", err)
				return
			}
			fmt.Printf("Sent: %s\n", message)
		}
	}
}

func main() {
	http.HandleFunc("/", handleConnections)

	fmt.Println("Server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
