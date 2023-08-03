package funcs

import (
	"log"
	"net"
)

var conn_type = "tcp"

func CreateServerPort(port string) {
	listener, err := net.Listen(conn_type, port)
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
