package core

import (
	"github.com/gorilla/websocket"
	"github.com/libp2p/go-libp2p/core/network"
)

var (
	upgrader = websocket.Upgrader{}
)

const protocolID = "/chat/1.0.0"

type Peers map[network.Stream]bool
type IDMapT map[string]bool
