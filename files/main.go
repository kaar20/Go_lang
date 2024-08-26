package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	fmt.Println("Welecome To the World of Files...")

	files, err := os.Create("./file.txt")
	checkError(err)

	content := "This needs to be a file and a directory and not a directory"

	length, err := io.WriteString(files, content)
	checkError(err)
	fmt.Println("The Length of the File is : ", length)
	readFile("./file.txt")

	defer files.Close()

}

func checkError(err error) {
	if err != nil {
		panic(err)

	}

}

func readFile(filename string) {
	readedFile, err := ioutil.ReadFile(filename)
	checkError(err)

	fmt.Println(string(readedFile))

}
