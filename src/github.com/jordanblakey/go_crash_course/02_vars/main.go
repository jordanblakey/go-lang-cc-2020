package main

import "fmt"

// MAIN VAR TYPES
// string
// bool
// int(8|16|32|64)
// uint(8|16|32|64|ptr)
// byte - alias for uint8
// rune - alias for int32
// float(32|64)
// complex(64|128)

// Can instantiate + assign vars and consts outside func bodies
// Cannot reassign outside func bodies, though
var exampleVar string = "Example"

const exampleConst int = 42

func main() {
	// Using var
	var name string = "Jordan"
	var age int = 34
	fmt.Println(name, age)

	// fmt.Printf accepts a range of format "verbs"
	// %v - default format
	// %#v - Go-syntax representation of the value
	// %T - Go-syntax rep of the type
	// %% - Literal percentage sign; consumes no value
	// %t - Boolean
	// %b - Base 2 | %c Character (Unicode rep) | %d - Base 10
	// %f - Float
	// and so on...

	fmt.Printf("%T, %v\n", name, name)
	fmt.Printf("%T, %v\n", age, age)

	var isCool bool = true
	fmt.Printf("%T, %v\n", isCool, isCool)
	isCool = false
	fmt.Println("isCool:", isCool)

	const isHuman bool = true
	fmt.Printf("isHuman: %v, Type: %T\n", isHuman, isHuman)

	altInstantiation := "Some string."
	fmt.Printf("altInstantiation: %T, %v\n", altInstantiation, altInstantiation)

	// Var assignment shorthand
	name2 := "Brad" // has inferred type of string
	fmt.Printf("name2: %T, %v\n", name2, name2)
	size := 1.3 // Has inferred type of float64
	fmt.Printf("size: %T, %v\n", size, size)
	var f32 float32 = 1.1234 // type cast to float32 vs inferred type of float64
	fmt.Printf("f32: %T, %v\n", f32, f32)

	// Multiple var assignment shorthand
	name3, email := "Jose", "jose@gmail.com"
	fmt.Println(name3, email)
}
