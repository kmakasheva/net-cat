package main

import (
	"fmt"
	"net-cat/funcs"
	"os"
)

func main() {
	port := ":8080"
	if len(os.Args) > 2 {
		fmt.Println("[USAGE]: To start the TCP Chat server, run: ./TCPChat %s\n", port)
		return
	} else if len(os.Args) == 2 {
		port = ":" + os.Args[1]
	}
	funcs.StartTCPServer(port)
}
