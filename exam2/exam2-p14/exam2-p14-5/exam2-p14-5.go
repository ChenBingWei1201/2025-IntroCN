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
	message, err := reader.ReadString('\n')
	check(err)
	fmt.Printf("%s", message)

	writer := bufio.NewWriter(conn)
	newline := fmt.Sprintf("Welcome to This Number is Closer! I have a true value in 1-100 for you to figure out. Give me two numbers (in two lines). I will show which one is closer to the true value.\n")
	_, err = writer.WriteString(newline)
	check(err)
	writer.Flush()

	// Read second message (number)
	message, err = reader.ReadString('\n')
	check(err)
	fmt.Printf("%s", message)
	var first int
	fmt.Sscanf(message, "%d", &first)

	// Read third message (number)
	message, err = reader.ReadString('\n')
	check(err)
	fmt.Printf("%s", message)
	var second int
	fmt.Sscanf(message, "%d", &second)

	trueValue := 62

	var response string
	if first == trueValue && second == trueValue {
		response = "Bingo\n"
	} else {
		diff1 := abs(first - trueValue)
		diff2 := abs(second - trueValue)

		if diff1 == diff2 {
			response = "Equally close\n"
		} else if diff1 < diff2 {
			response = fmt.Sprintf("%d is closer\n", first)
		} else {
			response = fmt.Sprintf("%d is closer\n", second)
		}
	}

	_, err = writer.WriteString(response)
	check(err)
	writer.Flush()
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
