package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {
	fmt.Println("Structs in Golang!")

	advait := User{"Advait","advait.trivedi@google.com",true,21}
	fmt.Println(advait)
	fmt.Printf("Details for Advait are: %+v\n",advait)
	fmt.Printf("advait.Name: %v\n",advait.Name)
	fmt.Printf("advait.Email: %v\n",advait.Email)
	fmt.Printf("advait.Status: %v\n",advait.Status)
	fmt.Printf("advait.Age: %v\n",advait.Age)
}
