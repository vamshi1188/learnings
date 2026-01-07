package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	s, err := net.Dial("tcp", "127.0.0.1:2121")
	if err != nil {
		fmt.Println("failed to connect server", err)
	}
	defer s.Close()

	for {
		WriteToServer(s)
	}
}
func WriteToServer(conn net.Conn) {
	// defer conn.Close()

	read := bufio.NewReader(os.Stdin)

	bytes, err := read.ReadBytes(byte('\n'))
	if err != nil {
		fmt.Println("failed to read the bytes", err)
	}

	text := string(bytes)
	fmt.Fprintf(conn, text)

}
