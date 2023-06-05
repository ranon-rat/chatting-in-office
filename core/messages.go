package core

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"

	"github.com/libp2p/go-libp2p/core/network"
)

type Message struct {
	Author  string `json:"author"`
	Content string `json:"content"`
}

func WriteMessage() {

	for {
		msg := <-msgChan
		for p := range peers {
			if err := json.NewEncoder(p).Encode(msg); err != nil {
				delete(peers, p)
				p.Close()

			}
		}
	}
}

func readMessage(p network.Stream, d bufio.ReadWriter) {
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
		content, _ := DecryptMessage(msg.Content)
		fmt.Printf("\r%s > %s\n\r> ", msg.Author, content)
	}
}
func DecodeMessage(content string) {
	base64.NewEncoder(base64.StdEncoding, &bufio.Writer{}).Write([]byte("abc"))
}
func WriteMSG(content string) {
	var msg Message

	msg.Content, _ = EncryptMessage(content[:len(content)-1])
	msg.Author = Author
	msgChan <- msg

}

func WritingMessage(author string) {
	Author = author
	for {
		fmt.Print("> ")
		reader := bufio.NewReader(os.Stdin)
		content, _ := reader.ReadString('\n')
		WriteMSG(content)
	}
}
