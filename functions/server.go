package netcat

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func Server(ip string) {
	dflt := "8080"
	if ip == "" {
		ip = dflt
	}
	ln, err := net.Listen("tcp", ":"+ip)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Listening on port" + ip)
	conn, err := ln.Accept()
	if err != nil {
		log.Fatal(err)
	}
	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print("Message Received:", string(message))
		newmessage, err := bufio.NewReader(os.Stdin).ReadString('\n')
		conn.Write([]byte(newmessage + "\n"))
	}
}
