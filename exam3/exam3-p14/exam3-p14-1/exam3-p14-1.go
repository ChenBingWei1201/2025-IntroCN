package main

import "crypto/tls"

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
}
