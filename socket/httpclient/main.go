package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "httpbin.org:80")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Println("connected!")

	_, err = conn.Write([]byte("GET /get HTTP/1.0\r\n\r\n"))
	if err != nil {
		panic(err)
	}
	data := make([]byte, 1000)
	n, err := conn.Read(data)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data[:n]))
}
