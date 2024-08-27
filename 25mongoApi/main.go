package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kaar20/mongoapi/router"
)

func main() {
	fmt.Println("Server Is Running")
	r := router.Router()
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening")
}
