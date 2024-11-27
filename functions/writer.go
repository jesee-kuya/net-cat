package netcat

import (
	"bufio"
	"log"
	"net"
	"os"
)

func Writer(conn net.Conn, channel chan string) {
	for {
		newmessage, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		conn.Write([]byte(newmessage + "\n"))
		channel <- newmessage
	}
}
