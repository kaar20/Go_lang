package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	welcome := "Welcome to the User Input Section"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Fadlan Gali Waxaa Rabto: ")

	input, _ := reader.ReadString('\n')
	fmt.Println("The input of the user is :", input)

	fmt.Printf("The output type of the user is %T :",input)

}
