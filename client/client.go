package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	const serverAddress string = ":3005"

	conn, err := net.Dial("tcp", serverAddress)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fileName := os.Args[1]

	_, err = conn.Write([]byte(fileName))
	if err != nil {
		panic(err)
	}

	buffer := make([]byte, 2048)

	n, err := conn.Read(buffer)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Bytes received: %d\n%s\n", n, string(buffer))
}
