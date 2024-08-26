package main

import "fmt"

func main() {
	// Defer works like Lifo(Last in First Out)

	defer fmt.Println("Hello")

	fmt.Println("Im first one Displyed while Defer is in the top")

	defer fmt.Println("First Defer ")

}
