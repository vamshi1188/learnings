package main

import (
	"fmt"
	"net"
)

func main() {

	listner, err := net.Listen("tcp", "localhost:2121")
	if err != nil {

		fmt.Println("failed to establish connection", err)
	}

	defer listner.Close()
}
