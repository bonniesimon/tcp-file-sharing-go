package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	const address string = ":3005"
	const CHUNK_SIZE int = 2048

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

		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		check(err)

		fileName := string(buffer[:n])
		fmt.Println("Read filename from client: ", fileName)

		// fileInfo, _ := os.Stat(string(fileName))

		// numberOfChunks := fileInfo.Size() / int64(CHUNK_SIZE)

		buffer = make([]byte, CHUNK_SIZE)

		file, err := os.OpenFile(string(fileName), os.O_RDWR, 0755)
		check(err)

		for {
			n, err := file.Read(buffer)
			if err != nil {
				if err == io.EOF {
					break
				}
				panic(err)
			}

			fmt.Printf("============================Buffer==============================\n%s\n=======================END==================\n", string(buffer[:n]))
			fmt.Printf("Characters read: %d\n\n\n\n\n", n)

			_, err = conn.Write(buffer[:n])
			check(err)
			break // We break here because I'm testing by sending only the first chunk of data
		}

		fmt.Printf("read the following data: %s\n", string(fileName))
	}
}
