package util

import (
	"github.com/gorilla/websocket"
	"log"
)

func ListenToMessages(conn *websocket.Conn, ch chan []byte) {
	// Spin up thread for intercepting feed url messages
	go func() {
		defer func(conn *websocket.Conn, feedCh chan []byte) {
			err := conn.Close()
			if err != nil {
				log.Fatalln("Error closing feed connection: ", err)
			}

			close(feedCh)
		}(conn, ch)

		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Fatalln("Error receiving feed message:", err)
				return
			}
			ch <- message
		}
	}()

}
