package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	cert, err := tls.LoadX509KeyPair("client.cer", "client.key")
	check(err)
	// skip checking the certificate
	config := tls.Config{Certificates: []tls.Certificate{cert}, InsecureSkipVerify: true}
	conn, _ := tls.Dial("tcp", "140.112.41.208:12000", &config)
	defer conn.Close()

	writer := bufio.NewWriter(conn)
	_, errw := writer.WriteString("PLAY\n")
	check(errw)
	writer.Flush()

	scanner := bufio.NewScanner(conn)
	if scanner.Scan() {
		fmt.Printf("%s\n", scanner.Text())
	}
}
