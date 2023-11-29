package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to video on slices!")

	var fruitList = []string{"Apple", "Banana", "Mango"}
	fmt.Printf("Type of fruitList is %T\n", fruitList)

	fruitList = append(fruitList, "Tomato", "Peach")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:])
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	fruitList = append(fruitList[:3])
	fmt.Println(fruitList)

	highScores := make([]int, 4)
	highScores[0] = 234
	highScores[1] = 235
	highScores[2] = 267
	highScores[3] = 239

	highScores = append(highScores, 555, 666, 321)
	fmt.Println(highScores)

	fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	sort.Strings(fruitList)
	fmt.Println(fruitList)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))
	var myFloats = make([]float64, 4)
	myFloats[0] = 1.2
	myFloats[1] = 2.2
	myFloats[2] = 0.2
	myFloats[3] = 3.1
	fmt.Println(myFloats)
	sort.Float64s(myFloats)
	fmt.Println(myFloats)

	// how to remove a value from from slices based on index

	var courses = []string{"reactjs", "nodejs", "golang", "solidity", "javascript", "python"}
	fmt.Println(courses)
	var index int = 3	// remove solidity
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
