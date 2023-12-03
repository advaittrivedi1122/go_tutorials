package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Printf("\n\n\t\t!! Welcome to Dice Game in Golang !!\n\n")

	rand.Seed(time.Now().UnixNano())

	var result string = ""
	var canMove bool = false
	var hasSecondChance bool = false
	var currentInput string = ""
	reader := bufio.NewReader(os.Stdin)

	for currentInput != "end" {

		if !hasSecondChance {
			fmt.Println("----------------------------------------------------------------------------")
			fmt.Printf("\n\nEnter your choice (press enter to start or continue, end to quit) : ")
			input, _ := reader.ReadString('\n')
			currentInput = strings.TrimSpace(input)
		}

		if currentInput == "end" {
			fmt.Println("\n\nThanks for playing... Bye!\n")
			break
		}

		dice := rand.Intn(6) + 1
		fmt.Println("\nDice :", dice)

		switch dice {
			case 1:
				hasSecondChance = false
				if !canMove {
					result = "Oops! You cannot move yet."
				} else {
					result = "You need to move 1 tile"
				}
				fmt.Println("Result :", result)
				break
			case 2:
				hasSecondChance = false
				if !canMove {
					result = "Oops! You cannot move yet."
				} else {
					result = "You need to move 2 tiles"
				}
				fmt.Println("Result :", result)
				break
			case 3:
				hasSecondChance = false
				if !canMove {
					result = "Oops! You cannot move yet."
				} else {
					result = "You need to move 3 tiles"
				}
				fmt.Println("Result :", result)
				break
			case 4:
				hasSecondChance = false
				if !canMove {
					result = "Oops! You cannot move yet."
				} else {
					result = "You need to move 4 tiles"
				}
				fmt.Println("Result :", result)
				break
			case 5:
				hasSecondChance = false
				if !canMove {
					result = "Oops! You cannot move yet."
				} else {
					result = "You need to move 5 tiles"
				}
				fmt.Println("Result :", result)
				break
			case 6:
			hasSecondChance = true
			if !canMove {
				result = "Congratulations! You can start playing now. Play dice again"
				} else {
					result = "Woohoo! You got six!! Move 6 tiles and play dice again."
				}
				fmt.Println("Result :", result)
				canMove = true
			break
		}
	}

}
