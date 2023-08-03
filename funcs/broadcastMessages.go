package funcs

import "fmt"

func Broadcaster() {
	for {
		select {
		case msg := <-joinChannel :
			clientMutex.Lock()
			for _, client := range clients {
				if msg.name == client.name {
					for _, w := range messageHistory{
						fmt.Fprintf(client.conn, "%s", "\r")
						fmt.Fprintf(client.conn, "%s[%s][%s]:", w, msg.time, client.name)
					}
				}
				if msg.name != client.name {
					fmt.Fprintf(client.conn, "\n%s %s\n[%s][%s]:", msg.name, msg.text, msg.time, client.name)
				}
			}
			clientMutex.Unlock()
		case msg := <-messageChannel:
			clientMutex.Lock()
			for _, client := range clients {
				if msg.name != client.name {
					fmt.Fprintf(client.conn, "\r[%s][%s]:\r", msg.time, msg.name)
					fmt.Fprintf(client.conn, "[%s][%s]:%s\n", msg.time, msg.name, msg.text)
				}
				fmt.Fprintf(client.conn, "[%s][%s]:", msg.time, client.name)
			}
			clientMutex.Unlock()
		case msg := <-leaveChannel :
			clientMutex.Lock()
			for _, client := range clients {
				if client.name != msg.name {
					fmt.Fprintf(client.conn, "\n%s %s\n[%s][%s]:", msg.name, msg.text, msg.time, client.name)
				}
			}
			clientMutex.Unlock()
		}
	}
}