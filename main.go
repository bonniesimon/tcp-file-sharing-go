package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	const address string = ":3005"

	listener, err := net.Listen("tcp", address)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
	defer listener.Close()

	for {
		fmt.Printf("Server ready to accept connections to %s\n", address)

		conn, err := listener.Accept()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
			os.Exit(1)
		}

		fmt.Printf("Connected to %s\n", address)

		buf := make([]byte, 256)

		n, _ := conn.Read(buf)

		fmt.Printf("Read %d bytes\n", n)
		fmt.Printf("read the following data: %s\n", string(buf))
	}
}
