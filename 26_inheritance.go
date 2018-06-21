package main

import "fmt"

/*
	Go does not support inheritance, however it does support composition. The generic definition of composition is "put together".
	One example of composition is a car. A car is composed of wheels, engine and various other parts.
*/

type author struct {
	firstName string
	lastName string
	bio string
}

func (a author) fullName() string {
	return fmt.Sprintf("%s %s", a.firstName, a.lastName)
}

type post struct {
	title string
	content string
	author			// anonymous field
}

func (p post) details() {
	fmt.Println("Title: ", p.title)
	fmt.Println("Content: ", p.content)
	fmt.Println("Author: ", p.fullName()) //Whenever one struct field is embedded in another, Go gives us the option
											 // to access the embedded fields as if they were part of the outer struct.
	fmt.Println("Bio: ", p.author.bio)
}

func main() {
	author := author{"krishna", "chaitanya", "Geek"}

	post := post{"Inheritance in Go", "Go supports composition instead of inheritance", author}

	post.details()
}