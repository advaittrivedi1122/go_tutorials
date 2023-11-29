package main

import "fmt"

func main() {
	fmt.Println("Maps in Golang!")

	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["PY"] = "Python"
	languages["GO"] = "Golang"
	languages["SOL"] = "Solidity"

	fmt.Printf("Type of all languages: %T\n", languages)
	fmt.Println("List of all languages:", languages)

	fmt.Println("JS shorts for:", languages["JS"])
	fmt.Println("PY shorts for:", languages["PY"])
	fmt.Println("GO shorts for:", languages["GO"])
	fmt.Println("SOL shorts for:", languages["SOL"])
	
	delete(languages, "SOL")
	fmt.Println("SOL deleted:", languages["SOL"])
	fmt.Println(languages["SOL"] == "")		// true

	// loops are interesting in golang

	for key, value := range languages {
		fmt.Printf("Key: %v, Value: %v\n", key, value)
	}
}