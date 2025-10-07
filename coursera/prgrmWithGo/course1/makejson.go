package main

import (
	"encoding/json"
	"fmt"
)

func Jsonfunc() {

	var name string
	var addr string

	fmt.Print("enter your name: ")
	fmt.Scan(&name)
	fmt.Print("enter your addr: ")
	fmt.Scan(&addr)

	details := map[string]string{

		"name":    name,
		"address": addr,
	}

	data, err := json.Marshal(details)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}
