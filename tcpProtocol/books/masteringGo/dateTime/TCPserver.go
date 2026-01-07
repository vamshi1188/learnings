package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	Arguments := os.Args

	if len(Arguments) == 1 {
		fmt.Println("please provide port number")
		return
	}
	port := ":" + Arguments[1]
	l, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println(err)
	}
	defer l.Close()

	c, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		if strings.TrimSpace(string(netData)) == "STOP" {
			fmt.Println("Exiting TCP server")
			return
		}

		fmt.Println(string(netData))

		t := time.Now()
		mytime := t.Format(time.RFC3339)
		c.Write([]byte(mytime))

	}

}
