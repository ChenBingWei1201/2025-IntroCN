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
	conn, errc := net.Dial("tcp", "140.112.41.208:11991")
	check(errc)
	defer conn.Close()
}
