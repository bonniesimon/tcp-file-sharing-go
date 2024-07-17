package main

import (
	"net"
)

func main() {
	const serverAddress string = "127.0.0.1:3005"

	conn, _ := net.Dial("tcp", serverAddress)
	defer conn.Close()

	var msg []byte = []byte{'c', 'a', 'k', 'e'}

	conn.Write(msg)
}
