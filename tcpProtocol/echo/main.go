package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("usage : go run main.go <port>")
		os.Exit(1)
	}

	port := fmt.Sprintf(":%s", os.Args[1])

	listener, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("failed to create listener", err)
		os.Exit(1)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("failed to accept connection", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	for {
		bytes, err := reader.ReadBytes(byte('\n'))
		if err != nil {
			if err != io.EOF {
				fmt.Println("failed to read data ", err)
			}
			return
		}
		fmt.Printf("request: %s", bytes)

		line := fmt.Sprintf("Echo: %s", bytes)
		fmt.Printf("response: %s", line)

		_, err = conn.Write([]byte(line))

		if err != nil {
			fmt.Println("failed to send the data", err)
		}
	}
}
