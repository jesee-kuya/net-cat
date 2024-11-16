package main

import (
	"os"

	netcat "netcat/functions"
)

func main() {
	args := os.Args
	if args[1] == "-s" {
		netcat.Server(args[2])
	} else {
		netcat.Client(args[2])
	}
}
