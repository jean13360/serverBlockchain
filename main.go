// You can edit this code!
// Click here and start typing.
package main

import (
	"blockchain/blockchain"
	"blockchain/websocket"
	"log"
	"net/http"
)

func main() {
	router := NewRouter()
	go blockchain.Init()
	go PeerToPeer.InitServer()
	//	PairToPair.SendMessage("message Envoyé")
	log.Fatal(http.ListenAndServe(":2012", router))
}
