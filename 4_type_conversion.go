package main

import "fmt"

func main() {

	// Go is very strict about explicit typing. There is no automatic type promotion or conversion.
	a := 10
	b := 11.65

	fmt.Println(a + int(b))  // We need to convert type before doing operations
}
