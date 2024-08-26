package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "https://simad.edu.so/"

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	// fmt.Printf("Response type is  %T", response)
	readResponse, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	// fmt.Println(string(readResponse))
	content := string(readResponse)
	fmt.Printf("Content is :", content)

	defer response.Body.Close()

}
