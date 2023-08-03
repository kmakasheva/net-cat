package funcs

import (
	"log"
	"net"
)

var connection_type = "tcp"

func CreateServerPort(port string) {
<<<<<<< HEAD
	ls, err := net.Listen(connection_type, port)
=======
	listener, err := net.Listen(conn_type, port)
>>>>>>> 1e16a9ee378faadc932e5488d612c0e5a9195bba
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
