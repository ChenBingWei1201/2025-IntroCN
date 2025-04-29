package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Launching server...")
	ln, _ := net.Listen("tcp", ":12000")
	defer ln.Close()

	for {
		conn, _ := ln.Accept()
		defer conn.Close()
	}
}
