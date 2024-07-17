package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	address := "127.0.0.1:3005"
	listener, err := net.Listen("tcp", address)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
	defer listener.Close()

	conn, err := listener.Accept()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	buf := make([]byte, 256)

	n, _ := conn.Read(buf)

	fmt.Printf("Read %d bytes\n", n)
	fmt.Printf("read the following data: %s\n", string(buf))
}
