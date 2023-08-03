package funcs

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

func HandleConnection(conn net.Conn) {
	timeClient := time.Now().Format("2006-01-02 15:04:05")
	Welcome(conn)
	nickname := GetNickname(conn)
	tempClient := Client{
		name: nickname,
		addr: conn.RemoteAddr().String(),
		conn: conn,
	}
	clientMutex.Lock()
	clients[nickname] = tempClient
	if len(clients) > 10 {
		clientMutex.Lock()
		tempClient.conn.Write([]byte("Error: Chat is full!"))
		delete(clients, nickname)
		tempClient.conn.Close()
		clientMutex.Unlock()
		return
	}
	clientMutex.Unlock()
	clientMutex.Lock()
	joinChannel <- NewMessage("has entered the chat...", conn, tempClient, timeClient)
	clientMutex.Unlock()
	fmt.Fprintf(conn, "[%s][%s]:", timeClient, nickname)
	input := bufio.NewScanner(conn)

	for input.Scan() {
		time := time.Now().Format("2006-01-02 15:04:05")
		if input.Text() == "" {
			fmt.Fprintf(conn, "Error: Message cannot be blank.\n")
			fmt.Fprintf(conn, "[%s][%s]:", time, nickname)
			continue
		}
		if IsCorrectInput(input.Text(), conn, time, nickname) == false {
			fmt.Fprintf(conn, "Error: Incorrect input.\n")
			fmt.Fprintf(conn, "[%s][%s]:", time, nickname)
			continue
		}
		text := fmt.Sprintf("[%s][%s]:%s\n", time, nickname, input.Text())
		clientMutex.Lock()
		messageHistory = append(messageHistory, text)
		messageChannel <- NewMessage(input.Text(), conn, tempClient, time)
		clientMutex.Unlock()
	}
	clientMutex.Lock()
	delete(clients, nickname)
	leaveChannel <- NewMessage("has left our chat...", conn, tempClient, timeClient)
	conn.Close()
	clientMutex.Unlock()
}
