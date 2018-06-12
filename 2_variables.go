package main

import "fmt"

// VARIABLE SYNTAX: var name type
// Default value assigned is 0

func main()  {
	var age int8
	fmt.Println("My age is: ", age)

	age = 25
	fmt.Println("My age is: ", age)

	// If a variable has an inital value, specifying type is not required
	var newAge = 24
	fmt.Println("My age is: ", new_age)

	// Multiple variables
	var length, width = 16, 17
	fmt.Println("width is", width, "length is", length)

	// Shorthand
	name, age := "naveen", 29
	fmt.Println("my name is", name, "age is", age)
}
