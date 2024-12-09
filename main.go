package main

import (
	"fmt"
	"net/http"
	"time"
)

func ping(url string) {
	for {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Failed to ping %s: %v\n", url, err)
		} else {
			fmt.Printf("Pinged %s: %s\n", url, resp.Status)

			resp.Body.Close()
		}

		time.Sleep(1 * time.Minute)
	}
}

func main() {
	url := "https://www.google.com"
	fmt.Printf("Starting ping service for %s\n", url)
	go ping(url)

	select {}
}