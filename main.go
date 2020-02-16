package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("Request failed")

type response struct {
	url    string
	status string
}

func main() {
	results := make(map[string]string)
	c := make(chan response)

	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
	}

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
func hitUrl(url string, c chan<- response) {
	res, err := http.Get(url)
	status := "OK"
	if err != nil || res.StatusCode >= 400 {
		status = "FAIELD"
	}
	c <- response{url: url, status: status}
}
