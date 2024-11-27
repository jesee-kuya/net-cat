package netcat

import (
	"fmt"
	"log"
	"net"
)

// type Client struct {
// 	name string
// 	conn net.Conn
// }

func Server(ip string) {
	var channel chan string = make(chan string)
	var msg string
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
		go Reader(conn, channel)
		go Writer(conn, channel)

		for {
			msg += <-channel
			fmt.Print(msg)
		}
	}
}
