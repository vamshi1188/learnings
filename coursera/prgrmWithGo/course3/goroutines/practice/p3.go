package main

import (
	"fmt"
	"runtime"
)

func GorunMax() {
	fmt.Println("cpus: ", runtime.NumCPU())
	runtime.GOMAXPROCS(2)
	fmt.Println(runtime.NumCgoCall())
	fmt.Println(runtime.Version())
}
