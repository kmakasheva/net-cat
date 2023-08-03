package funcs

import (
	"net"
	"sync"
)

var (
	clientMutex       sync.Mutex
	clients     = make(map[string]Client)
	historytext = []string{}
	leaving     = make(chan Message)
	messageChannel   = make(chan Message)
	joinChannel        = make(chan Message)
)

type Client struct {
	
	name string
	// айпи адрес пользователя
	addr string
	// интерфейс конекшна
	conn net.Conn
}

// cтруктура сообщения
type Message struct {
	// текст сообщения
	text string
	// айпи адрес пользователя
	address string
	// имя пользователя
	name string
	// время
	time string
	// история сообщении
	history []string
}
