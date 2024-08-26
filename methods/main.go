package main

import "fmt"

func main() {
	fmt.Println("Welcome to the Struct Module ")

	kaar := Student{"Mohamed Abdiaziz Ali", 23, 99.99, true}
	fmt.Printf("kaar is: %+v\n", kaar)

	kaar.studentAge()
	// Note : No inheritance in golang ; No Super or parent
}

type Student struct {
	Name        string
	Age         int
	Grade       float64
	IsGraduated bool
}

func (s Student) studentAge() {
	s.Age = 99
	fmt.Println("Student age: ", s.Age)

}
