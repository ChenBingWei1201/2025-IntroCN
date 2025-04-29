package main

import (
	"fmt"
	"bufio"
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

	writer := bufio.NewWriter(conn)
	_, errw := writer.WriteString("PLAY\n")
	check(errw)
	writer.Flush()

	scanner := bufio.NewScanner(conn)
	if scanner.Scan() {
		fmt.Printf("%s\n", scanner.Text())
	}
	
	// Get user's guess
	var character rune
	fmt.Println("Guess a character:")
	fmt.Scanf("%c", &character)

	// Send the guessed character to server
	_, errw = writer.WriteString(fmt.Sprintf("%c\n", character))
	check(errw)
	writer.Flush()

	// Read the server's response
	if scanner.Scan() {
		fmt.Printf("%s\n", scanner.Text())
	}
}
