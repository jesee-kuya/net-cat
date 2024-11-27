package netcat

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func Reader(conn net.Conn) {
	message, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("Message Received:", string(message))
	newmessage, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	conn.Write([]byte(newmessage + "\n"))
}
