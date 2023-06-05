package core

import (
	"bufio"
	"context"

	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/protocol"
)

func ConnectMDNS(node host.Host, peerChan chan peer.AddrInfo, protocolID protocol.ID) {
	for {
		peerAddrInfo := <-peerChan
		if e := IDmap[peerAddrInfo.ID.String()]; e {
			continue
		}
		// Connect to the node at the given address.
		if err := node.Connect(context.Background(), peerAddrInfo); err != nil {
			panic(err)
		}
		// con esto solo me conecto a ese puerto
		s, _ := node.NewStream(context.Background(), peerAddrInfo.ID, protocolID)
		if e := peers[s]; !e {
			WriteMSG("hello im new here :D ")

		}
		peers[s] = true
		IDmap[s.ID()] = true
		go readMessage(s, *bufio.NewReadWriter(bufio.NewReader(s), bufio.NewWriter(s)))

	}
}
func NewConnection(p network.Stream) {
	//	go readMessage(p, *bufio.NewReadWriter(bufio.NewReader(p), bufio.NewWriter(p)))
	IDmap[p.ID()] = true

	peers[p] = true

}
