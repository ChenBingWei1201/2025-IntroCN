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

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	reader := bufio.NewReader(conn)

	// Read first message (PLAY)
	message, errr := reader.ReadString('\n')
	check(errr)
	fmt.Printf("%s", message)

	writer := bufio.NewWriter(conn)
	newline := fmt.Sprintf("Welcome to This Number is Closer! I have a true value in 1-100 for you to figure out. Give me two numbers (in two lines). I will show which one is closer to the true value.\n")
	_, errw := writer.WriteString(newline)
	check(errw)
	writer.Flush()

	// Read second message (number)
	message, errr = reader.ReadString('\n')
	check(errr)
	fmt.Printf("%s", message)

	// Read third message (number)
	message, errr = reader.ReadString('\n')
	check(errr)
	fmt.Printf("%s", message)
}
