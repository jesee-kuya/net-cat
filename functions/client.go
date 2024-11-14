package netcat

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func Client(ip string) {
	conn, err := net.Dial("tcp", ip)
	if err != nil {
		log.Fatal(err)
	}
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(conn, text+"\n")
		msg, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server: " + msg)
	}
}
