package netcat

import (
	"fmt"
	"net"
)

type Client struct {
	name string
	conn net.Conn
}

var Clients []Client

func Server(port string) {
	if port == "" {
		port = "8989"
	}
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Listening on the port :" + port)

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go Reader(conn)
	}
}
