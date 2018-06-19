package main

import "fmt"

func find(num int, nums ...int)  {
	fmt.Printf("type of nums is %T\n", nums) //type of nums is []int
	found := false
	for i, v := range nums {
		if v == num {
			fmt.Println(num, "found at index", i, "in", nums)
			found = true
		}
	}
	if !found {
		fmt.Println(num, "not found in ", nums)
	}
	fmt.Printf("\n")
}

func change(s ...string) {
	s[0] = "Go"
}

func main() {
	/*
	A variadic function is a function that can accept variable number of arguments. If the last parameter of a function
	is denoted by ...T, then the function can accept any number of arguments of type T for the last parameter.Please note
	that only the last parameter of a function is allowed to be variadic.
	 */
	 find(89, 89, 90, 91) // nums are converted by the compiler to a slice of type int []int{89, 90, 95}

	 find(89, []int{1,2,3}) // Won't work as variadic argument requires numbers and it internally coverts to a slice,
	 // As we are passing slice complier tries to convert slice to slice an it throws error

	 // TO MAKE THE ABOVE WORK
	 welcome := []string{"hello", "world"}

	 change(welcome...) // Convert slice to variadic

	 // Change will change the first element in welcome as it is passed as a reference

}