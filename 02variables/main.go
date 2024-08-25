package main

import "fmt"

func main() {
	// var Mohamed string="Mohamed Abdiaziz Ali is a Muslim ,Mumin Ku kalsoon Rabigiis";
	// fmt.Println(Mohamed)
	// fmt.Printf("Length of the string is: %T \n",Mohamed)
	// Declaring without var and type
	guled := "Mohamed Abdiaziz Ali is a Muslim ,Mumin Ku Kalsoon Rabigiis"
	fmt.Println(guled)
	fmt.Printf("Length of the string is: %d \n", len(guled))
	// Declare with var and type
	//     var myString string = "Mohamed Abdiaziz Ali is a Muslim ,Mumin Ku Kalsoon Rabigiis"
	//     fmt.Println(myString)
	//     fmt.Printf("Length of the string is: %d \n",len(myString))

	var a string
	var b int
	var c bool
	fmt.Println(a, b, c)

	a = "abcdefghijklmnopqr"

	fmt.Println(a)
}
