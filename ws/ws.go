package ws

import (
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Handler implements http.Handler to handle *websocket.Conn
// For example:
//
//	h := func(conn *ws.Conn) {
//		for {
//			messageType, message, err := conn.ReadMessage()
//			if err != nil {
//				log.Printf("Failed to read message: %v", err)
//				break
//			}
//
//			log.Printf("Received message: %s", message)
//
//			err = conn.WriteMessage(messageType, message)
//			if err != nil {
//				log.Printf("Failed to write message: %v", err)
//				break
//				}
//			}
//		}
//	}
//
// route.Route("/ws_echo", ws.Handler(h))
func Handler(handle func(conn *websocket.Conn)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer conn.Close()
		handle(conn)
	})
}
