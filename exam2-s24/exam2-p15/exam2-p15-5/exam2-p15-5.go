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
	ln, _ := net.Listen("tcp", ":12000")
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
	newline := fmt.Sprintf("Welcome to Hangman Not Quite! I have an English word for you to figure out. Give me a character. I will show where the character appears in the word.\n")
	_, errw := writer.WriteString(newline)
	check(errw)
	writer.Flush()

	// Read second message (character)
	message, errr = reader.ReadString('\n')
	check(errr)
	fmt.Printf("%s", message)

	// Send response for the character
	newline = fmt.Sprintf("Server got %s\n", message)
	_, errw = writer.WriteString(newline)
	check(errw)
	writer.Flush()
}
