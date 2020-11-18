package main

import "fmt"

func main() {
	// Basic For Loop
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}

	// Shorthand
	for i := 1; i <= 10; i++ {
		fmt.Printf("Number: %d\n", i)
	}

	// FizzBuzz
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			if i%5 == 0 {
				fmt.Println(fmt.Sprint(i) + ": FizzBuzz")
			} else {
				fmt.Println(fmt.Sprint(i) + ": Fizz")
			}
		} else if i%5 == 0 {
			fmt.Println(fmt.Sprint(i) + ": Buzz")
		} else {
			fmt.Println(i)
		}
	}

}
