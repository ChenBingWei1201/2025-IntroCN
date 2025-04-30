package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Launching server...")
	ln, _ := net.Listen("tcp", "140.112.41.208:30202")
	defer ln.Close()

	for {
		conn, _ := ln.Accept()
		defer conn.Close()
	}
}
