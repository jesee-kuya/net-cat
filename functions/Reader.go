package netcat

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func Reader(conn net.Conn) {
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
		conn.Write([]byte("Wrong name"))
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
