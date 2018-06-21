package main

import "fmt"

/*
	there are some situations where the program cannot simply continue executing after an abnormal situation. In this
	case we use panic to terminate the program. When a function encounters a panic, its execution is stopped, any deferred
	functions are executed and then the control returns to its caller. This process continues until all the functions of
	the current goroutine have returned at which point the program prints the panic message, followed by the stack trace and then terminates.

	It is possible to regain control of a panicking program using recover which we will discuss later in this tutorial.

	panic and recover can be considered similar to try-catch-finally idiom in other languages except that it is rarely
	used and when used is more elegant and results in clean code.

	One important factor is that you should avoid panic and recover and use errors where ever possible. Only in cases
	where the program just cannot continue execution should a panic and recover mechanism be used.
*/

/*
	func panic(interface{})

	The argument passed to panic will be printed when the program terminates.

*/

func fullName(firstName *string, lastName *string) {

	/*
		When a panic is encountered, the program execution terminates, prints the argument passed to panic followed by the stack trace.
	 */
	// If a deferred function call is present, it is executed and then the control returns to its caller.
	defer fmt.Println("deferred call in fullName")
	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")

}

func main() {
	firstName := "Elon"
	fullName(&firstName, nil)
	fmt.Println("returned normally from main")
}