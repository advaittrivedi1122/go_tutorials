package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=1s0iawmndnodrew34"

func main() {
	fmt.Println("Welcome to handling urls in Golang!")
	fmt.Println(myUrl)

	// Parsing
	result, _ := url.Parse(myUrl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Port())
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)

	queryParams := result.Query()
	fmt.Printf("Type of Query Params are %T\n", queryParams)
	fmt.Println("Coursename is:", queryParams.Get("coursename"))	// "reactjs"
	fmt.Println("Coursename is:", queryParams["coursename"])		// ["reactjs"]

	for _, value := range(queryParams) {
		fmt.Println("Param is: ", value)
	}

	// always pass reference of url.URL
	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "/learn",
		RawQuery: "user=advait",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println("Another URL: ", anotherURL)
}
