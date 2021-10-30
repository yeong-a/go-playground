package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":8000")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Println("Server connected!")

	data := make([]byte, 100)
	n, err := conn.Read(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(n, "bytes read!")
	fmt.Println(string(data))
}
