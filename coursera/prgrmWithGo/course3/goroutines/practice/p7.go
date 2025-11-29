package main

import (
	"fmt"
	"time"
)

func FetchUrl(url string) {

	fmt.Println("fetching ", url)

	time.Sleep(time.Millisecond * 500)

	fmt.Println("done", url)
}

func FetchingUrl() {

	urls := []string{"a.com", "b.com", "c.com"}

	for _, i := range urls {

		go FetchUrl(i)
	}
	time.Sleep(time.Second)
}
