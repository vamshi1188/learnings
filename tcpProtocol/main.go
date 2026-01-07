package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {

	listener, err := net.Listen("tcp", ":2121")
	if err != nil {
		fmt.Println("failed to open tcp connection", err)
		return
	}

	defer listener.Close()

	for {
		c, err := listener.Accept()

		if err != nil {
			fmt.Println("failed to accept the incoming connection", err)
			return
		}
		fmt.Println(c.RemoteAddr())
		ReadData(c)

	}

}

func ReadData(conn net.Conn) {
	defer conn.Close()

	for {
		reader := bufio.NewReader(conn)

		text, _ := reader.ReadBytes(byte('\n'))

		fmt.Print(string(text))
	}
}
