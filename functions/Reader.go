package netcat

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func Reader(conn net.Conn) {
	defer conn.Close()
	if len(Clients) == 4 {
		conn.Write([]byte("[Sorry the chat if full, try again later]"))
		return
	}
	var client Client
	client.conn = conn
	entry, err := Readfile("./src/logo.txt")
	if err != nil {
		conn.Write([]byte(fmt.Sprintf("%v", err)))
		return
	}
	conn.Write([]byte(entry))
	name, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		conn.Write([]byte("Connection Denied"))
		return
	}
	client.name = name
	Clients = append(Clients, client)

	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		if message != "" {
			Chat(client.conn, message)
		}
		
	}
}
