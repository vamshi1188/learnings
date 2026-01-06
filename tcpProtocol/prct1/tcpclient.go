package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	arguments := os.Args

	if len(arguments) == 1 {
		fmt.Println("provide port Usage: go run main.go <port>")
		return
	}
	connect := arguments[1]

	c, err := net.Dial("tcp", connect)

	if err != nil {
		fmt.Println("failed to connect: ", connect)
		return
	}
	defer c.Close()

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(c, text)
		if strings.TrimSpace(text) == "stop" {
			fmt.Println("client exiting.....")
			return
		}
	}

}
