package main

import (
	crand "crypto/rand"
	"flag"
	"fmt"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/ranon-rat/chatting-in-office/core"
)

const protocolID = "/chat/1.0.0"

func main() {
	fmt.Println(`
                                 
	 _           _       _       _   
	| |_ ___ _ _| |_ ___| |_ ___| |_ 
	| . |  _| | |   |  _|   | .'|  _|
	|___|_| |___|_|_|___|_|_|__,|_|  
									 
	
	

 simple chatting tool for using it in the office
 

	`)

	author := flag.String("username", "anon", "")
	rendezvous := flag.String("channel", "public", "")
	flag.Parse()

	r := crand.Reader
	// Creates a new RSA key pair for this host.
	prvKey, _, err := crypto.GenerateKeyPairWithReader(crypto.RSA, 2048, r)
	if err != nil {
		panic(err)
	}
	node, err := libp2p.New(
		libp2p.ListenAddrStrings("/ip4/127.0.0.1/tcp/0"),
		libp2p.Identity(prvKey),
	)
	if err != nil {
		panic(err)
	}
	defer node.Close()

	// con esto lo que hago es simplemente dejar que en un puerto me pueda comunicar por medio de un protocolo
	node.SetStreamHandler(protocolID, core.NewConnection)

	peerChan := core.InitMDNS(node, *rendezvous)
	core.InitCipher(*rendezvous)
	go core.ConnectMDNS(node, peerChan, protocolID)
	go core.WriteMessage()
	core.WritingMessage(*author)

}
