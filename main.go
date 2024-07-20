package main

import (
	"fmt"
	"io"
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

		received_data, err := io.ReadAll(conn)
		if err != nil {
			panic(err)
		}

		err = os.WriteFile("tmp/receive/file.txt", received_data, 0644)
		if err != nil {
			panic(err)
		}

		fmt.Printf("read the following data: %s\n", string(received_data))
		fmt.Printf("Wrote to file tmp/receive/file.txt\n")
	}
}
