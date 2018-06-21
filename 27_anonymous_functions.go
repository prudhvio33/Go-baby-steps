package main

import "fmt"

func main(){
	a := func() string{ return "Hello"}  // Type of it is func()

	fmt.Println(a())

	b := func(a int, b int) int {return a + b}

	fmt.Println(b(1,3))
}
