package main

import (
	"net"
	"os"
)

func main() {
	const serverAddress string = ":3005"

	conn, _ := net.Dial("tcp", serverAddress)
	defer conn.Close()

	data, err := os.ReadFile("tmp/send/file.txt")
	if err != nil {
		panic(err)
	}

	_, err = conn.Write(data)

	if err != nil {
		panic(err)
	}
}
