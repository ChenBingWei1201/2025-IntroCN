package main

import (
	"bufio"
	"net"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	conn, errc := net.Dial("tcp", "140.112.41.208:12000")
	check(errc)
	defer conn.Close()

	writer := bufio.NewWriter(conn)
	_, errw := writer.WriteString("PLAY\n")
	check(errw)
	writer.Flush()
}
