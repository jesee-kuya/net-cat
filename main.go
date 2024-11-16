package main

import (
	"os"

	netcat "netcat/functions"
)

func main() {
	args := os.Args
	if args[1] == "-s" {
			netcat.Server()
	} else {
			netcat.Client(args[2])
	}
}
