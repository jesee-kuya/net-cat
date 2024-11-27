package netcat

import (
	"bufio"
	"log"
	"net"
)

func Reader(conn net.Conn, channel chan string) {
	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		channel <- message
	}
}
