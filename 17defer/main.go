package main

import "fmt"

func main() {
	defer fmt.Println("First defer")	// gets executed in the end
	defer fmt.Println("Second defer")	// gets executed in the end
	fmt.Println("Welcome to defer statement in Golang!")
	fmt.Println("Gets printed after defer.")
	myDefer()
}

func myDefer(){
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
