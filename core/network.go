package core

import (
	"bufio"
	"context"
	"crypto/cipher"

	"github.com/gorilla/websocket"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/protocol"
)

func ConnectMDNS(node host.Host, peerChan chan peer.AddrInfo, protocolID protocol.ID, msgChan chan Message, author string, peers Peers, IDmap IDMapT, ws *websocket.Conn, key cipher.Block) {
	for {
		peerAddrInfo := <-peerChan
		if e := IDmap[peerAddrInfo.ID.String()]; e {
			continue
		}
		// Connect to the node at the given address.
		if err := node.Connect(context.Background(), peerAddrInfo); err != nil {
			return
		}
		// con esto solo me conecto a ese puerto
		s, _ := node.NewStream(context.Background(), peerAddrInfo.ID, protocolID)
		if e := peers[s]; !e {
			WriteMSG("Buenas mis compañeros del MKUltra me acabo de conectar :D ", msgChan, author, key)

		}
		peers[s] = true
		IDmap[s.ID()] = true
		go readMessage(s, *bufio.NewReadWriter(bufio.NewReader(s), bufio.NewWriter(s)), peers, IDmap, ws, key)

	}
}
