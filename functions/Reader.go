package netcat

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func Reader(conn net.Conn) {
	message, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("Message Received:", string(message))
}
