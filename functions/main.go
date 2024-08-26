package main

import "fmt"

func main() {
	fmt.Println("Welcome to the Functions Module")

fmt.Println(	add(1, 3))
	fmt.Println(	values(1, 3, 4, 5, 6, 6, 7, 7))

}

func add(i int, y int) int {
	return i + y

}

func values(values ...int) int {
	result := 0

	for _, value := range values {
		result += value
		// fmt.Println(value)
	}

	return result
}
