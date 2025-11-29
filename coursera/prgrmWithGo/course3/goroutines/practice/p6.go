package main

import (
	"fmt"
	"time"
)

func calc(n int) {
	sum := 0
	for i := 0; i < 1e7; i++ {
		sum += i
	}
	fmt.Println("Task", n, "done")
}

func CountCalc() {
	for i := 0; i < 4; i++ {
		go calc(i)
	}
	time.Sleep(time.Second)
}
