package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("The Time Now is")
	present := time.Now()

	fmt.Println("The Time Now is: ", present)
	fmt.Println(present.Format("2006-01-02  Monday  15:04:05"))
}
