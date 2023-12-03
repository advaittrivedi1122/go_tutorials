package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Switch case in Golang!")

	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Enter your number : ")
	input, _ := reader.ReadString('\n')
	
	num, _ := strconv.ParseInt(strings.TrimSpace(input), 10, 64)

	switch num {
		case 1:
			fmt.Println("Num is", num, "and case is 1")		// without break, it still will not go in blocks below
		case 2:
			fmt.Println("Num is", num, "and case is 2")
			fallthrough						// Telling Lexer explicitly to go in the block below
		case 3:
			fmt.Println("Num is", num, "and case is 3")
			fallthrough
		case 4:
			fmt.Println("Num is", num, "and case is 4")
			break
		case 5:
			fmt.Println("Num is", num, "and case is 5")
		default:
			fmt.Println("Default ...")
	}

}
