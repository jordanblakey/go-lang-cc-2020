package main

import "fmt"

func main() {
	// Arrays
	var fruitArray [2]string

	// Assign values
	fruitArray[0] = "Apple"
	fruitArray[1] = "Orange"

	fmt.Println(fruitArray)
	fmt.Println(fruitArray[0])
	fmt.Println(fruitArray[1])

	// Declare and assign
	vegetableArray := []string{"Tomato", "Cucumber", "Carrot"}
	fmt.Println(vegetableArray)
	fmt.Println(vegetableArray[1:2])
	fmt.Println(vegetableArray[1:])
	fmt.Println(vegetableArray[:3])
	fmt.Println(vegetableArray[0:])
}
