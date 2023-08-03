package funcs

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func GetNickname(conn net.Conn) string {
	conn.Write([]byte("[Enter your name]:"))
	data, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return ""
	}
	temp := strings.TrimSpace(string(data))

	if temp == "" || len(temp) == 0 || len(temp) > 20 {
		fmt.Fprintln(conn, "Error: Incorrect input")
		return GetNickname(conn)
	}

	for _, i := range temp {
		if i < 32 || i > 127 {
			fmt.Fprintln(conn, "Error: Incorrect input\n")
			return GetNickname(conn)
		}
	}
	for client := range clients {
		if client == temp {
			fmt.Fprintf(conn, "User already exists\n")
			return GetNickname(conn)
		}
	}
	return temp
}

func IsCorrectInput(s string, conn net.Conn, time string, username string) bool {
	for _, w := range s {
		if w < 32 || w > 127 {
			return false
		}
	}
	return true
}

func Welcome(conn net.Conn) {
	file, err := os.ReadFile("logo.txt")
	if err != nil {
		fmt.Printf("Couldn't read this file")
	}
	imagetxt := string(file)
	conn.Write([]byte("Welcome to our Chat Room!\n" + imagetxt + "\n"))
}

func NewMessage(msg string, conn net.Conn, cl Client, time string) Message {
	return Message{
		text:    msg,
		address: cl.addr,
		name:    cl.name,
		time:    time,
		history: messageHistory,
	}
}
