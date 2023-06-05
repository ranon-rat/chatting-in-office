package core

import (
	crand "crypto/rand"
	"net/http"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/network"
)

func JoiningChannel(w http.ResponseWriter, r *http.Request) {
	ws, _ := upgrader.Upgrade(w, r, nil)
	defer ws.Close()
	author := r.URL.Query()["username"][0]
	channel := r.URL.Query()["channel"][0]

	msgChan := make(chan Message)
	key, _ := InitCipher(channel)
	ra := crand.Reader
	// Creates a new RSA key pair for this host.
	prvKey, _, err := crypto.GenerateKeyPairWithReader(crypto.RSA, 2048, ra)
	if err != nil {
		return
	}
	node, err := libp2p.New(
		libp2p.ListenAddrStrings("/ip4/127.0.0.1/tcp/0"),
		libp2p.Identity(prvKey),
	)
	if err != nil {
		return
	}
	defer node.Close()

	peers := make(Peers)
	IDmap := make(IDMapT)
	// con esto lo que hago es simplemente dejar que en un puerto me pueda comunicar por medio de un protocolo
	node.SetStreamHandler(protocolID, func(s network.Stream) {
		peers[s] = true
		IDmap[s.ID()] = true

	})
	peerChan := InitMDNS(node, channel)
	go ConnectMDNS(node, peerChan, protocolID, msgChan, author, peers, IDmap, ws, key)
	go WriteMessage(msgChan, peers, IDmap)
	for {
		var msg Message

		if err := ws.ReadJSON(&msg); err != nil {
			return
		}
		WriteMSG(msg.Content, msgChan, author, key)
	}
}
