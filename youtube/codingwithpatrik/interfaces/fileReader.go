package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func FileReader() {
	file, _ := os.Open("data.txt")
	resp, _ := http.Get("http://example.com")
	buffer := strings.NewReader("2024-01-10 09:15:33 ERROR db connection failed host=localhost:5432")

	fmt.Println("---FIle---")
	io.Copy(os.Stdout, file)
	fmt.Println("\n---http Body ---")
	io.Copy(os.Stdout, resp.Body)
	fmt.Println("\n---String---")
	io.Copy(os.Stdout, buffer)
}
