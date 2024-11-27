package netcat

import (
	"fmt"
	"log"
	"net"
)

func Server(ip string) {
	if ip == "" {
		ip = "8989"
	}
	ln, err := net.Listen("tcp", ":"+ip)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Listening on the port " + ip)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go Reader(conn)
		go Writer(conn)
	}
}
