package main

import (
	"fmt"
	"time"
)

func PrintCOncurently() {
	go fmt.Println("Hello from goroutine")

	fmt.Println("Hello from main")
	time.Sleep(time.Second * 1)
}
