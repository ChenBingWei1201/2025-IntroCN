package main

import (
	"net"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	conn, errc := net.Dial("tcp", "127.0.0.1:12000")
	check(errc)
	defer conn.Close()
}
