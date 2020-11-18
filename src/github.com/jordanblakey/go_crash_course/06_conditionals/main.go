package main

import "fmt"

func idColor(color string) {
	if color == "red" {
		fmt.Println("color is red.")
	} else if color == "blue" {
		fmt.Println("color is blue")
	} else {
		fmt.Println("color is NOT blue or red")
	}
}

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


}