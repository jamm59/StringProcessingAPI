package handler

import (
	"fmt"
	"log"
	"net/http"

	websocket "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Handler() {

	http.HandleFunc("/todo", func(w http.ResponseWriter, r *http.Request) {
		// Upgrade upgrades the HTTP server connection to the WebSocket protocol.
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Print("upgrade failed: ", err)
			return
		}
		defer conn.Close()

		// Continuosly read and write message
		for {
			mt, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("read failed:", err)
				break
			}
			fmt.Println(string(rune(mt)))
			var output string = string(message)
			message = []byte(output)
			err = conn.WriteMessage(mt, message)
			if err != nil {
				log.Println("write failed:", err)
				break
			}
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>Hello from Go!</h1>")
	})

	http.ListenAndServe(":8080", nil)
}
