package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in Golang!")

	content := "This needs to go in a file."
	
	file, err := os.Create("./myGoFile.txt")
	if err != nil {
		panic(err)
		// panic :- shut down execution of program and show the error
	}

	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}

	fmt.Println("length is:", length)

	defer file.Close()

	readFile("./myGoFile.txt")
}

func readFile(filename string) {
	dataByte, err := ioutil.ReadFile(filename)
	checkNilError(err)
	fmt.Println("Text Data inside the file is :-\n", string(dataByte))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}