package main

import (
	"fmt"
	"net"
)

func main() {

	port := ":2121"

	l, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	for {
		k, err := l.Accept()
		if err != nil {
			fmt.Println(err)
		}
		go Acceptccept(k)

	}

}

func Acceptccept(conn net.Conn) {
	fmt.Println(conn)
}
