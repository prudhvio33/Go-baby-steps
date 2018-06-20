package main

import "fmt"

/*
	A structure is a user defined type which represents a collection of fields. It can be used in places where it makes
	sense to group the data into a single unit rather than maintaining each of them as separate types.

	For instance a employee has a firstName, lastName and age. It makes sense to group these three properties into a
	single structure employee.
*/

type Employee struct {
	fname string
	lname string
	age   int
}

type EmployeeV2 struct {
	firstName, lastName string
	age, salary         int
}

// snippet creates a anonymous structure employee.
var employee struct {
	firstName, lastName string
	age int
}

// NESTED STRUCTS
type Address struct {
	city string
}

type Information struct{
	name string
	address Address
}

func main()  {
	emp1 := Employee{
		fname: "Krishna",
		lname: "Chaitanya",
		age:   24,
	}

	emp2 := EmployeeV2{"krishna", "teja", 23, 50000}

	fmt.Println("Employee 1", emp1)
	fmt.Println("Employee 2", emp2)

	// Accessing values in a structure

	fmt.Println(emp1.fname)

	emp8 := &EmployeeV2{"Sam", "Anderson", 55, 6000}
	fmt.Println("First Name:", emp8.firstName)
	fmt.Println("Age:", emp8.age)

	var info Information
	info.name = "chaitanya"
	info.address = Address{
		city: "Chicago",
	}

	fmt.Println(info)



}