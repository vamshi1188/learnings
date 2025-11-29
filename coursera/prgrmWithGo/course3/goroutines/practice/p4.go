package main

import (
	"fmt"
	"time"
)

func Work() {

	fmt.Println(" gorouitine")
	time.Sleep(time.Second)
}
func MeasureTime() {
	start := time.Now()

	for i := 1; i <= 5; i++ {
		go Work()
	}

	time.Sleep(time.Second * 2)
	fmt.Println("elapsed time : ", time.Since(start))
}
