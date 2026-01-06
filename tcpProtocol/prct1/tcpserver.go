package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func main() {

	port := "localhost:2121"
	listner, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Printf("failed to open connection on%v \n", port, err)
		return
	}

	defer listner.Close()

	for {

		conn, err := listner.Accept()
		if err != nil {
			fmt.Println("failed to accept connection", err)
		}

		ClientAccept(conn)
	}
}

func ClientAccept(conn net.Conn) {

	defer conn.Close()

	reader := bufio.NewReader(conn)

	for {
		bytes, err := reader.ReadBytes(byte('\n'))
		if err != nil {
			if err != io.EOF {
				fmt.Println("failed to read the data")
			}
		}

		fmt.Print(string(bytes))

	}

}
