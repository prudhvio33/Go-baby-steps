package main

import "fmt"

func main()  {
	var a [3]int // array of length 3
	fmt.Println(a) //[0,0,0]

	a[0] = 100

	b := [...]int{12, 78, 50} // ... makes the compiler determine the length
	fmt.Println(b)

	c := [3]int{5, 78, 8}
	var d [5]int
	c = d //not possible since [3]int and [5]int are distinct types

	/*
		Arrays in Go are value types and not reference types. This means that when they are assigned to a new variable,
	a copy of the original array is assigned to the new variable. If changes are made to the new variable, it will not
	be reflected in the original array. Similarly when arrays are passed to functions as parameters, they are passed by
	value and the original array in unchanged.

	 */
	e := [...]string{"usa", "canada", "china"}
	f := e
	fmt.Println(e) // ["usa", "canada", "china"]
	f[0] = "Singapore"
	fmt.Println(f) // ["Singapore", "canada", "china"]

	fmt.Println("Length of array ", len(a))

	for i, v := range a {//range returns both the index and value
		fmt.Printf("%d the element of a is %.2f\n", i, v)
	}


}
