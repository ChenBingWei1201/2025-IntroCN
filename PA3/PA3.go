package main

import (
	"fmt"
	"bufio"
	"os"
	"net"
	"io"
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

	inputName := ""
	fmt.Printf("Input filename: ")
	fmt.Scanf("%s", &inputName)

	input, inputErr := os.Open(inputName)
	check(inputErr)
	defer input.Close()

	fileInfo, fileInfoErr := input.Stat()
	check(fileInfoErr)
	fileSize := fileInfo.Size()

	writer := bufio.NewWriter(conn)
	_, errw := writer.WriteString(fmt.Sprintf("%d\n", fileSize))
	check(errw)
	writer.Flush()

	_, errCopy := io.Copy(writer, input)
	check(errCopy)
	writer.Flush()
	
	scanner := bufio.NewScanner(conn)
	if scanner.Scan() {
		fmt.Printf("Send the file size first: %d\n", fileSize)
		fmt.Printf("Server says: %s\n", scanner.Text())
	}
}
