package main

import "fmt"

func main() {

	// var ptr *int
	// fmt.Println("ptr is :", ptr)

	myNumber := 23

	var ptr = &myNumber

	fmt.Println("the position of myNumber value is :", ptr)
	fmt.Println("the value of myNumber is :", *ptr)

	*ptr = 45

	fmt.Println("the value of myNumber after changing it through pointer is :", myNumber)
	// var ptr
}
