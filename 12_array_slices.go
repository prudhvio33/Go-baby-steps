package main

import "fmt"

func main() {
	/*
		A slice does not own any data of its own. It is just a representation of the underlying array. Any modifications
	done to the slice will be reflected in the underlying array.
	 */
	var a = [5]int{1, 2, 3, 4, 5}
	var b = a[1:3]

	fmt.Println(b)

	/*
		length: Length of the current slice
		capacity: Number of elements in the underlying array starting from the slice creation to end of underlying array
	 */

	fmt.Println(len(b)) // 2
	fmt.Println(cap(b)) // 4

	/*
		APPENDING TO SLICE
		------------------
		As we already know arrays are restricted to fixed length and their length cannot be increased. Slices are dynamic
		and new elements can be appended to the slice using append function. The definition of append function is func
		append(s []T, x ...T) []T. x ...T in the function definition means that the function accepts variable number of
		arguments for the parameter x. These type of functions are called variadic functions.

		If slices are backed by arrays and arrays themselves are of fixed length then how come a slice is of dynamic
		length. Well what happens under the hoods is, when new elements are appended to the slice, a new array is created.
		The elements of the existing array are copied to this new array and a new slice reference for this new array is
		returned. The capacity of the new slice is now twice that of the old slice.
	 */

	cars := []string{"Ferrari", "Honda", "Ford"}
	fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars)) //capacity of cars is 3
	cars = append(cars, "Toyota")
	fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars)) //capacity of cars is doubled to 6

	/*
		APPENDING TO A NIL
		------------------
		The zero value of a slice type is nil. A nil slice has length and capacity 0. It is possible to append values
	to a nil slice using the append function
	 */

	var names []string
	if names == nil {
		names = append(names, "krishna", "chaitanya")

	}
	fmt.Println(names)
	names = append(names,"teja")
	fmt.Println(names)

}
