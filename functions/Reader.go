package netcat

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
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
		conn.Write([]byte("Connection Denied"))
		return
	}
	name = strings.TrimSpace(name)
	client.name = name
	Clients = append(Clients, client)

	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			Clients = Remove(Clients, conn)
			conn.Close()
			break
		}
		t := fmt.Sprintf("[%v][%v]: %v", time.Now().Format("2006-01-02 15:04:05"), client.name, message)
		Chat(client.conn, t)
	}
}
