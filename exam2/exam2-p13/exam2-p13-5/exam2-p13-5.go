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
	conn, errc := net.Dial("tcp", "140.112.41.208:12000")
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
	num := 0
	fmt.Println("Guess the 1st numbers:")
	fmt.Scanf("%d", &num)

	// Send the guessed num to server
	_, errw = writer.WriteString(fmt.Sprintf("%d\n", num))
	check(errw)
	writer.Flush()

	fmt.Println("Guess the 2nd numbers:")
	fmt.Scanf("%d", &num)

	// Send the guessed num to server
	_, errw = writer.WriteString(fmt.Sprintf("%d\n", num))
	check(errw)
	writer.Flush()

	// Read the server's response
	if scanner.Scan() {
		fmt.Printf("%s\n", scanner.Text())
	}
}
