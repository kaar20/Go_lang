package main

import "fmt"

func main() {
	fmt.Println("Welcome to the Array Module")

	var fruitList [4]string

	fruitList[0] = "Mohamed"
	fruitList[1] = "Abdiaziz"
	fruitList[2] = "Ali"

	var fruitList1 = [5]string{
		"Apple",
		"Orange",
		"Banana",
		"Grape",
		"Mango",
	}

	fmt.Println("fruitList is : ", fruitList1)
}
