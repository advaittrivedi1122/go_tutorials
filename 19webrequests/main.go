package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("Welcome to webRequests in Golang!")

	response, err := http.Get(url)
	checkNilError(err)
	
	fmt.Printf("Response is of type: %T\n", response)

	defer response.Body.Close()
	// caller's responsibilty to close the connection

	dataBytes, err := ioutil.ReadAll(response.Body)
	checkNilError(err)
	content := string(dataBytes)
	fmt.Println("content is :", content)
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}