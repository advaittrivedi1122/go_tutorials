package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in Golang!")

	days := []string{
		"Sunday",
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
	}

	fmt.Println("\nDays in week are", days)
	fmt.Println()
	
	for i := 0; i < len(days) ; i++ {
		fmt.Println(days[i])
	}
	
	fmt.Println()
	for i := range(days) {
		fmt.Println("days in range :", days[i])
	}
	
	fmt.Println()
	for index, day := range(days) {
		fmt.Printf("index is %v and day is %v\n", index, day)
	}

	rogueValue := 1

	fmt.Println()
	for rogueValue <= 10 {
		if rogueValue == 5 {
			goto five	// also breaks the loop
		}
		if rogueValue == 3 {
			rogueValue++ // important to make rogueValue = 4 in order to continue further
			continue
		}
		if rogueValue == 6 {
			break
		}
		fmt.Println("rogueValue is", rogueValue)
		rogueValue++
	}

	five:
		fmt.Println("I was sent here by the for loop to inform you that loop is currently at index 5")
}
