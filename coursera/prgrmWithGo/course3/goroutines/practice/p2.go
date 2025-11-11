package main

import (
	"fmt"
	"time"
)

func MultipleTasks() {

	for i := 1; i <= 3; i++ {
		go Tasks(i)
	}
	time.Sleep(time.Second * 3)
}

func Tasks(id int) {
	fmt.Println("Task ", id, " started")
	time.Sleep(time.Second * 1)
	fmt.Println("Task ", id, " completed")
}
