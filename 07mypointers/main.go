package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on pointers!")
	myNumber := 23
	var ptr *int = &myNumber	// referencing
	var ptrToPtr **int = &ptr	// double pointer (pointer referencing to a pointer)

	fmt.Println()
	fmt.Println("myNumber =",myNumber)
	fmt.Println("&myNumber =",&myNumber)

	fmt.Println()
	fmt.Println("ptr =",ptr) // address of myNumber
	fmt.Println("&ptr =",&ptr) // address of pointer itself
	fmt.Println("*ptr =",*ptr)	// dereferencing (value of myNumber = 23)

	fmt.Println()
	fmt.Println("ptrToPtr =",ptrToPtr) // address of first pointer
	fmt.Println("&ptrToPtr =",&ptrToPtr) // address of double pointer itself
	fmt.Println("*ptrToPtr =",*ptrToPtr) // address of myNumber
	fmt.Println("**ptrToPtr =",**ptrToPtr)	// dereferencing double pointer (value of myNumber = 23)

	*ptr = *ptr + 2
	fmt.Println()
	fmt.Println("My changed number is", myNumber)
}
