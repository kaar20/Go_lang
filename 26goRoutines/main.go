package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	websiteList := []string{
		"https://google.com",
		"https://simad.edu.so",
		"https://fb.com",
	}

	for _, website := range websiteList {

		go getStatusCode(website)
		wg.Add(1)
	}
	wg.Wait()
}

func getStatusCode(website string) {
	defer wg.Done()
	res, err := http.Get(website)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer res.Body.Close() // It's good practice to close the response body

	// Get the current time
	currentTime := time.Now().Format(time.RFC1123)

	// Print the status code with the current time
	fmt.Printf("%s: %d Status Code for %s\n", currentTime, res.StatusCode, website)
}
