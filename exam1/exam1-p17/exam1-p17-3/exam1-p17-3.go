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
	conn, errc := net.Dial("tcp", "140.112.41.208:11993")
	check(errc)
	defer conn.Close()

	num := 0
	fmt.Scanf("%d", &num)
	writer := bufio.NewWriter(conn)
	_, errw := writer.WriteString(fmt.Sprintf("%d\n", num))
	check(errw)
	writer.Flush()

	scanner := bufio.NewScanner(conn)

	if scanner.Scan() {
		fmt.Printf("%s\n", scanner.Text())
	}
}
