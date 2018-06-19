package main

import (
	"fmt"
	"unsafe"
)

/*
	bool

	Numeric Types
		int8, int16, int32, int64, int = -128 to 127, -32768 to 32767, -2147483648 to 2147483647, -9223372036854775808 to 9223372036854775807,
										 32 or 64 depending on arch

		uint8, uint16, uint32, uint64, uint = 0 to 255, 0 to 65535, 0 to 4294967295, 0 to 18446744073709551615
		float32, float64 = float64 is the default type
		complex64, complex128
		byte is an alias of uint8
		rune is an alias of int32

	string
*/

func main()  {
	a := true
	b := false

	var result = a && b

	fmt.Println(result)

	var c = 89 // If no variable is specified it will be inferred from the right hand side
	d := 95

	fmt.Println("value of a is", c, "and b is", d)
	fmt.Printf("type of a is %T, size of a is %d bytes", c, unsafe.Sizeof(c)) //type and size of a
	fmt.Printf("\ntype of b is %T, size of b is %d bytes", d, unsafe.Sizeof(d)) //type and size of b

	firstName := "krishna"
	fmt.Println("\nMy first name is %s", first_name)
}


