package services

import (
	"context"
	"fmt"

	"github.com/gorilla/websocket"
)

// WebsocketService :nodoc
type WebsocketService struct {
	connection     *websocket.Conn
	twitterService TwitterService
}

// NewWebsocketService :nodoc
func NewWebsocketService(connection *websocket.Conn, twitterService TwitterService) *WebsocketService {
	return &WebsocketService{
		connection:     connection,
		twitterService: twitterService,
	}
}

// Serve :nodoc
func (service *WebsocketService) Serve(ctx context.Context) {
	defer func() {
		service.connection.Close()
	}()

	go func() {
		for {
			select {
			case message := <-service.twitterService.TwitterChan:
				fmt.Println("receive message, will passing to websocket")

				w, err := service.connection.NextWriter(websocket.TextMessage)
				if err != nil {
					return
				}

				w.Write(message)

				fmt.Println("Send data to websocket connection...")

				if err = w.Close(); err != nil {
					return
				}
			}
		}
	}()

	fmt.Println("[*] Websocket connection is alive")
	forever := make(chan bool)
	<-forever
}
