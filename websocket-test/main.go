package main

import (
	"log"
	"net/http"

	"github.com/olahol/melody"
)

func main() {
	mel := melody.New()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		mel.HandleRequest(w, r)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	mel.HandleMessage(func(s *melody.Session, msg []byte) {
		mel.Broadcast(msg)
	})

	mel.HandleConnect(func(s *melody.Session) {
		log.Println("New Connection!")
	})

	mel.HandleDisconnect(func(s *melody.Session) {
		log.Println("Disconnected!")
	})

	http.ListenAndServe(":4479", nil)
}
