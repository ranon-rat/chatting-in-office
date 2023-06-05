package core

import (
	"bufio"
	"crypto/cipher"
	"encoding/json"

	"github.com/gorilla/websocket"
	"github.com/libp2p/go-libp2p/core/network"
)

type Message struct {
	Author  string `json:"author"`
	Content string `json:"content"`
}

func WriteMessage(msgChan chan Message, peers Peers, IDmap IDMapT) {

	for {
		msg := <-msgChan
		for p := range peers {
			if err := json.NewEncoder(p).Encode(msg); err != nil {
				delete(peers, p)
				delete(IDmap, p.ID())

				p.Close()

			}
		}
	}
}

func readMessage(p network.Stream, d bufio.ReadWriter, peers Peers, IDmap IDMapT, ws *websocket.Conn, key cipher.Block) {
	for {
		var msg Message
		// but at the same time it works as a io.Reader
		// hm
		err := json.NewDecoder(d).Decode(&msg)
		if err != nil {
			delete(peers, p)
			delete(IDmap, p.ID())
			p.Close()

			return
		}
		msg.Content, _ = DecryptMessage(msg.Content, key)
		ws.WriteJSON(msg)
	}
}

func WriteMSG(content string, msgChan chan Message, author string, key cipher.Block) {
	var msg Message

	msg.Content, _ = EncryptMessage(content, key)
	msg.Author = author
	msgChan <- msg

}
