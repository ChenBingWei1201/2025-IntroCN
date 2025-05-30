package main

import (
	"bufio"
	"fmt"
	"net"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("Launching server...")
	ln, _ := net.Listen("tcp", "140.112.41.208:30202")
	defer ln.Close()

	for {
		conn, _ := ln.Accept()
		defer conn.Close()

		reader := bufio.NewReader(conn)
		message, errr := reader.ReadString('\n')
		check(errr)
		fmt.Printf("%s", message)
	}
}
