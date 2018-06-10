package main

import "fmt"

/*
	for loop
		for initialisation; condition; post {}

	infinite loop
		for {}
 */
func main()  {
	for a:=0; a < 10; a++ {
		if a%2 == 0 {
			continue //All code present in a for loop after the continue statement will not be executed for the current iteration. The loop will move on to the next iteration.
		}
		fmt.Println(a)
	}
}