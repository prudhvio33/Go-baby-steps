package main

import "fmt"

type add func(a int, b int) int // creates a new function type add which accepts two integer arguments and returns a integer.

func main()  {
	var a add = func(a int, b int) int {
		return a + b
	}
	s := a(5, 6)
	fmt.Println("Sum", s)
}