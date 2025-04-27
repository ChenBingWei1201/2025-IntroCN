package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("Launching server...")
	ln, _ := net.Listen("tcp", ":12002")
	defer ln.Close()

	for {
		conn, _ := ln.Accept()

		reader := bufio.NewReader(conn)
		req, err := http.ReadRequest(reader)
		check(err)

		path := strings.TrimPrefix(req.URL.Path, "/")

		file, err := os.Open(path)
		defer file.Close()

		if err != nil {
			fmt.Println("File not found")
			conn.Close()
			continue
		}

		fileInfo, err := file.Stat()
		check(err)
		fileSize := fileInfo.Size()
		fmt.Printf("File size = %d\n", fileSize)

		conn.Close()
	}
}
