package main

import "fmt"

func main() {
	fmt.Println("If else in Golang!")

	loginCount := 12
	isLoggedIn := true
	var result string

	if !isLoggedIn {
		result = "not logged in."
	} else if isLoggedIn && loginCount > 10{
		result = "logged in. Server traffic high. Kicking you off!"
	} else {
		result = "logged in."
	}

	fmt.Println("You are", result)

	number := 12

	if number % 2 == 0 {
		fmt.Printf("number %v is even\n", number)
	} else {
		fmt.Printf("number %v is odd\n", number)
	}

	if num := 11; num < 10 {
		fmt.Println("num is less than 10")
	} else {
		fmt.Println("num is not less than 10")
	}

	// if err != nil {
		// if error is not nil, there is some value in err so process the error
	// }
}
