package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	Arguments := os.Args

	if len(Arguments) == 1 {
		fmt.Println("provide port number")
	}
	server := "localhost" + ":" + Arguments[1]

	s, err := net.ResolveTCPAddr("tcp", server)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(s)

}
