package main

import "fmt"

func main() {
	x := 5
	y := 10

	// Will execute this block
	if x < y {
		fmt.Printf("%d is less than %d.\n", x, y)
	}

	// Will not execute this block
	if y < x {
		fmt.Printf("%d is less than %d\n", y, x)
	}

	var color string = "red"
	idColor(color)
	color = "blue"
	idColor(color)
	color = "green"
	idColor(color)

	// Switch
	switch color {
	case "red":
		fmt.Println("color is red.")
	case "blue":
		fmt.Println("olor is blue.")
	default:
		fmt.Println("color is NOT red or blue.")
	}
}

// If, else if, else
func idColor(color string) {
	if color == "red" {
		fmt.Println("color is red.")
	} else if color == "blue" {
		fmt.Println("color is blue.")
	} else {
		fmt.Println("color is NOT red or blue.")
	}
}
