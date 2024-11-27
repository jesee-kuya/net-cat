package netcat

import (
	"bufio"
	"log"
	"net"
)

func Reader(conn net.Conn) {
	var client Client
	client.conn = conn
	conn.Write([]byte("\n[ENTER YOUR NAME]: "))
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
