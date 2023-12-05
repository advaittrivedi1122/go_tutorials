package main

import "fmt"

func main() {
	greeter()
	fmt.Println("Welcome to Functions in Golang")

	var result int = add(2, add(2,3))
	fmt.Println("Result is:", result)

	allAdded, myMessage := addAll(1,2,3,4,5)
	fmt.Println("All numbers added is:", allAdded)
	fmt.Println("My message is:", myMessage)
}

func add(a int, b int) int {
	return a + b
}

func addAll(values ...int) (int,string) {
	total := 0
	for _, value := range(values) {
		total += value
	}
	return total, "return two values"
}

func greeter() {
	fmt.Println("Namaste from Golang!")
}