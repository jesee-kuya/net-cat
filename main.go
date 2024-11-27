package main

import (
	"os"

	netcat "netcat/functions"
)

func main() {
	args := os.Args
	var ip string
	if len(args) > 2 {
		return
	} else if len(args) == 2 {
		ip = args[1]
	}
	netcat.Server(ip)
}
