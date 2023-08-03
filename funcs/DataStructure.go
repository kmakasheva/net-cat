package funcs

import (
	"net"
	"sync"
)

var (
	clientMutex    sync.Mutex
	clients        = make(map[string]Client)
	messageHistory = []string{}
	leaveChannel   = make(chan Message)
	messageChannel = make(chan Message)
	joinChannel    = make(chan Message)
)

type Client struct {
	name string
	addr string
	conn net.Conn
}

type Message struct {
	text    string
	address string
	name    string
	time    string
	history []string
}
