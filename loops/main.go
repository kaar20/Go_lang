package main

import "fmt"

func main() {
	fmt.Println("Welcome to the Loops Module in Go")

	days := []string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"}

	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	for index, day := range days {
		fmt.Printf("Index: %d, Day: %s\n", index, day)
	}

}
