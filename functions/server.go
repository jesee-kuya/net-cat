package netcat

import (
	"fmt"
	"log"
	"net"
)

func Server(ip string) {
	dflt := "8989"
	if ip == "" {
		ip = dflt
	}
	ln, err := net.Listen("tcp", ":"+ip)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Listening on port" + ip)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go Reader(conn)
	}
}
