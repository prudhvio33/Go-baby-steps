package main

import "fmt"

/*
func functionname(parametername type) returntype {
 	//function body
}
 */
func add2Numbers(a int, b float32) int  {
	c := a + int(b)
	return c
}

// Multiple return values
func rectangleProps(a int, b int)(int, int) {
	return a, b
}

// Named return values
func rectProps(length, width float64)(area, perimeter float64) {
	area = length * width
	perimeter = (length + width) * 2
	return //no explicit return value
}

func main()  {
	fmt.Println(add2Numbers(10, 2.5))
	a, b := rectangleProps(10, 12)
	fmt.Println(a, b)
}
