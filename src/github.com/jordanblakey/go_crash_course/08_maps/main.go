package main

import "fmt"

func main() {
	// Instantiate map
	emails := make(map[string]string)

	// Add key value pairs
	emails["Bob"] = "bob@gmail.com"
	emails["Sharon"] = "sharon@gmail.com"
	emails["Alice"] = "alice@gmail.com"

	fmt.Println(emails)
	fmt.Println(emails["Bob"])

	// Get length of map
	fmt.Println(len(emails))

	// Delete map member
	delete(emails, "Bob")
	fmt.Println(emails)

	// Shorthand instantiation and key value pair creation
	emails2 := map[string]string{"Bob": "bob@gmail.com", "Sharon": "sharon@gmail.com"}
	emails2["Mike"] = "mike@gmail.com"
	fmt.Println(emails2)
}
