package urlcheck

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("Request failed")

type Response struct {
	url    string
	status string
}

func Check(urls []string) {
	results := make(map[string]string)
	c := make(chan Response)

	for _, url := range urls {
		go hitUrl(url, c)
	}

	fmt.Println("Waiting...")

	for i := 0; i < len(urls); i++ {
		result := <-c
		results[result.url] = result.status
	}

	for url, status := range results {
		fmt.Println(url, status)
	}
}

// c chan<-: c is a send-only channel (you cannot <- c)
func hitUrl(url string, c chan<- Response) {
	res, err := http.Get(url)
	status := "OK"
	if err != nil || res.StatusCode >= 400 {
		status = "FAIELD"
	}
	c <- Response{url: url, status: status}
}
