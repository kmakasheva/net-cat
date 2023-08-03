package funcs

import "fmt"

func MessageBroadcaster() {
	for {
		select {
		case joinMsg := <-joinChannel:
			HandleJoinMessage(joinMsg)

		case msg := <-messageChannel:
			HandleChatMessage(msg)

		case leaveMsg := <-leaveChannel:
			HandleLeaveMessage(leaveMsg)
		}
	}
}

func HandleJoinMessage(joinMsg Message) {
	clientMutex.Lock()
	for _, client := range clients {
		if joinMsg.name == client.name {
			for _, oldMessage := range messageHistory {
				fmt.Fprintf(client.conn, "\r%s [%s][%s]:", oldMessage, joinMsg.time, client.name)
			}
		} else {
			fmt.Fprintf(client.conn, "\n%s %s\n[%s][%s]:", joinMsg.name, joinMsg.text, joinMsg.time, client.name)
		}
	}
	clientMutex.Unlock()
}

func HandleChatMessage(msg Message) {
	clientMutex.Lock()
	for _, client := range clients {
		if msg.name != client.name {
			fmt.Fprintf(client.conn, "\r[%s][%s]:\r", msg.time, msg.name)
			fmt.Fprintf(client.conn, "[%s][%s]:%s\n", msg.time, msg.name, msg.text)
		}
		fmt.Fprintf(client.conn, "[%s][%s]:", msg.time, client.name)
	}
	clientMutex.Unlock()
}

func HandleLeaveMessage(leaveMsg Message) {
	clientMutex.Lock()
	for _, client := range clients {
		if client.name != leaveMsg.name {
			fmt.Fprintf(client.conn, "\n%s %s\n[%s][%s]:", leaveMsg.name, leaveMsg.text, leaveMsg.time, client.name)
		}
	}
	clientMutex.Unlock()
}
