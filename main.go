package main

import (
	"flag"
	"fmt"
)

func main() {
	l := flag.String("l", "8080",  "Listen mode, for inbound connects")
	flag.Parse()
	if *l != "" {
		fmt.Println(*l)
	}
}
