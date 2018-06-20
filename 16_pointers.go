package main

import "fmt"

/*
	A pointer is a variable which stores the memory address of another variable.

			a				b

			0x1203a124     156

							0x1203a124
 */
 func pass_address(val *int) {
 	*val++
 }

func modify(arr *[3]int)  {
	arr[0] = 90
}

func modifyV2(arr []int)  {
	arr[0] = 91
}

func main() {

	b := 255
	var a *int = &b // a will store address of b

	fmt.Printf("Type of a is %T\n", a)  // Type of a is *int
	fmt.Println("address of b is", a) 		// address of b is 0xc420084008
	fmt.Println("address of b is", &b) 		// address of b is 0xc420084008
	fmt.Println("value of a is", *a) 		// value of a is 255

	var c *int // This will store nil
	fmt.Println(c)

	// CHANGING THE VALUE IN B USING A
	*a++ // Take the value in that address and increment
	fmt.Println(b) //256
	pass_address(a)

	// PASS ARRAY TO A FUNCTION AND REFLECT CHANGES IN THE CALLER
	arr := [3]int{89, 90, 91}
	modify(&arr)
	fmt.Println(arr)

	modifyV2(arr[:])
	fmt.Println(arr)

	// Go does not support pointer arithmetic which is present in other languages like C.

	d := 100
	e := &d

	e++ //not possible
}
