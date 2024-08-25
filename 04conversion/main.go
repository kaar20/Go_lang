package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to the Sub Google Cloud Platform")
	fmt.Println("Please Rate Our life 1 to 8")

	rating := bufio.NewReader(os.Stdin)

	input, _ := rating.ReadString('\n')
	ratingConv, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println("Error parsing rating", err)
		// return
	} else {
		fmt.Println("Added Your Rating to  +10: ", ratingConv+10)
	}
	// if err !=nil{
	// 	fmt.Println("Error reading input", err)
	//     return
	// }
	// else{
	// 	fmt.Println("Added Your Rating to  +10: ",ratiratingConv + 10)
	// }
}
