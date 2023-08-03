package funcs

import (
	"log"
	"net"
)

var connection_type = "tcp"

func StartTCPServer(port string) {
	listener, err := net.Listen(connection_type, port)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Listening on the port: %s\n", port)
	go Broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go HandelConn(conn)
	}
}
