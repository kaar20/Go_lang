package main

import "fmt"

func main() {
	fmt.Println("Welcome to the Slices Module")

	highScore := make([]int, 4)
	highScore[0] = 100
	highScore[1] = 95
	highScore[2] = 88
	highScore[3] = 92
	// fmt.Println("High Score List is : ", highScore)
	highScore = append(highScore, 100, 100)
	// fmt.Println("After Appending new score : ", highScore)
	highScore = highScore[:3]
	// fmt.Println("After Slicing the list : ", highScore)
	highScore[2] = 98
	// fmt.Println("After Updating the score at index 2 : ", highScore)

	// Remove Value from the Slice
	fmt.Println(highScore)

	index := 2

	highScore = append(highScore[:index], highScore[index+1:]...)
	fmt.Println(highScore)

}
