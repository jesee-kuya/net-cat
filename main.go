package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
)

func main() {
	f := &flag.FlagSet{}

	var l string
	f.StringVar(&l, "l", "8080", "Listen mode, for inbound connects")
	buf := bytes.NewBuffer([]byte{})

	f.SetOutput(buf)
	err := f.Parse(os.Args[1:])

	if err == nil {
		fmt.Println(l)
	}else {
                fmt.Println("nc: missing port number")
                return
        }
}
