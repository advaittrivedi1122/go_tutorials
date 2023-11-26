package main

import "fmt"

// jwtToken := 300000 (walrus operator not allowed in global scope)
var jwtToken = 300000 // allowed

// first character capital -> public variable
const LoginToken string = "nxozidusdkasd"

func main() {
	var username string = "advait"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallValue uint8 = 255
	fmt.Println(smallValue)
	fmt.Printf("Variable is of type: %T \n", smallValue)

	var smallFloatValue float64 = 255.4552389230293434890
	fmt.Println(smallFloatValue)
	fmt.Printf("Variable is of type: %T \n", smallFloatValue)
	
	// default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)
	
	// implicit type (type inference)
	
	var website = "google.com"
	fmt.Println(website)
	
	// no var style (:= walrus operator)
	
	numberOfUsers := 300000
	fmt.Println(numberOfUsers)
	
	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)
}
