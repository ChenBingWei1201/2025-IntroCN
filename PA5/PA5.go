package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("Launching server...")
	ln, _ := net.Listen("tcp", "140.112.41.208:12002")
	defer ln.Close()

	for {
		conn, _ := ln.Accept()

		// Read received file size
		reader := bufio.NewReader(conn)
		message, errr := reader.ReadString('\n')
		check(errr)

		// Convert the message to integer
		sizeStr := strings.TrimSpace(message)
		fileSize, err := strconv.ParseInt(sizeStr, 10, 64)
		check(err)

		// Create the output file
		output, outputErr := os.Create("whatever.txt")
		check(outputErr)
		defer output.Close()

		// Create a limited reader to only read the expected number of bytes
		limitedReader := io.LimitReader(reader, fileSize)

		// Read the file content line by line
		lineReader := bufio.NewReader(limitedReader)
		writer := bufio.NewWriter(output)

		lineNum := 1
		bytesWritten := 0

		for {
			line, err := lineReader.ReadString('\n')
			if err != nil && err != io.EOF {
				check(err)
			}

			// Write line with line number
			if line != "" {
				lineWithNumber := fmt.Sprintf("%d %s", lineNum, line)
				bytesWritten += len(lineWithNumber)
				_, writeErr := writer.WriteString(lineWithNumber)
				check(writeErr)
				lineNum++
			}

			if err == io.EOF {
				break
			}
		}

		writer.Flush()

		// Print the upload file size and output file size
		fmt.Printf("Upload file size: %d\n", fileSize)
		fmt.Printf("Output file size: %d\n", bytesWritten)

		// Send acknowledgement back to client
		connWriter := bufio.NewWriter(conn)
		response := fmt.Sprintf("%d bytes received, %d bytes file generated\n", fileSize, bytesWritten)
		_, errw := connWriter.WriteString(response)
		check(errw)
		connWriter.Flush()

		conn.Close()
	}
}
