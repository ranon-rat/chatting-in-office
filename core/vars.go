package core

import (
	"crypto/cipher"

	"github.com/libp2p/go-libp2p/core/network"
)

var (
	peers   = make(map[network.Stream]bool)
	msgChan = make(chan Message)
	Caes    cipher.Block
	IDmap   = make(map[string]bool)
	Author  = ""
)
