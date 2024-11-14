package netcat

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func Server() {
	ln, err := net.Listen("tcp", "8080")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Listening on port 8080")
	conn, err := ln.Accept()
	if err != nil {
		log.Fatal(err)
	}
	for {
		msg, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print("Message Received:", string(msg))
		newmessage := strings.ToUpper(msg)
		conn.Write([]byte(newmessage + "\n"))

	}
}
