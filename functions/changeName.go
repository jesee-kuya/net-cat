package netcat

import (
	"bufio"
	"fmt"
	"strings"
)

func ChangeName(client Client) ([]Client, Client, error) {
	ext := fmt.Sprintf("%v has left the chat\n", client.name)
	client.conn.Write([]byte("[ENTER NEW NAME]: "))
	name, err := bufio.NewReader(client.conn).ReadString('\n')
	if err != nil {
		client.conn.Write([]byte("Connection Denied"))
		client.conn.Close()
		Chat(client.conn, ext)
		Clients = Remove(Clients, client.conn)
		return Clients, client, fmt.Errorf("connection closed")
	}
	name = strings.TrimSpace(name)

	for i := 0; i < len(Clients); i++ {
		if Clients[i].conn == client.conn {
			Clients[i].name = name
			client.name = name
		}
	}
	return Clients, client, nil
}
