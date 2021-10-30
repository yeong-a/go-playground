package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	conn, err := listener.Accept()
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	fmt.Println("Client connected!")

	n, err := conn.Write([]byte("Hello, client!"))
	if err != nil {
		panic(err)
	}
	fmt.Println(n, "bytes written!")
}
