package main

import (
	"net"
	"os"
)

func main() {
	const serverAddress string = ":3005"

	conn, _ := net.Dial("tcp", serverAddress)
	defer conn.Close()

	message := os.Args[1]

	conn.Write([]byte(message))
}
