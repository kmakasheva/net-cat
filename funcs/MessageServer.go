package funcs

import (
	"log"
	"net"
)

const connectionType = "tcp"

func StartTCPServer(port string) {
	listener, err := net.Listen(connectionType, port)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Listening on the port: %s\n", port)
	go MessageBroadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go HandleConnection(conn)
	}
}
