package main

import "fmt"

/*
	A method is just a function with a special receiver type that is written between the func keyword and the method name.
	The receiver can be either struct type or non struct type. The receiver is available for access inside the method.

	func (t Type) methodName(parameter list) {}

	1) Go is not a pure object oriented programming language and it does not support classes. Hence methods on types is
		a way to achieve behaviour similar to classes.
	2) Methods with same name can be defined on different types whereas functions with the same names are not allowed.

*/

type EmployeeV3 struct {
	name     string
	salary   int
	currency string
	age 	 int
}

/*
 displaySalary() method has Employee as the receiver type
*/
func (e EmployeeV3) displaySalary() {
	fmt.Printf("Salary of %s is %s%d", e.name, e.currency, e.salary)
}

func (e EmployeeV3) changeName(newName string) {
	e.name = newName
}

/*
Method with pointer receiver
*/
func (e *EmployeeV3) changeAge(newAge int) {
	e.age = newAge
}

func main() {
	emp1 := EmployeeV3 {
		name:     "Sam Adolf",
		salary:   5000,
		currency: "$",
		age: 25,
	}
	emp1.displaySalary() //Calling displaySalary() method of Employee type

	/*
		Pointer receivers vs value receivers
	----------------------------------------
		So far we have seen methods only with value receivers. It is possible to create methods with pointer receivers.
		The difference between value and pointer receiver is, changes made inside a method with a pointer receiver is
		visible to the caller whereas this is not the case in value receiver.
	 */

	 /*
	 	When to use pointer receiver and when to use value receiver
	 	-----------------------------------------------------------
		Generally pointer receivers can be used when changes made to the receiver inside the method should be visible to the caller.

		Pointers receivers can also be used in places where its expensive to copy a data structure. Consider a struct
	 	which has many fields. Using this struct as a value receiver in a method will need the entire struct to be copied
	 	which will be expensive. In this case if a pointer receiver is used, the struct will not be copied and only a
	 	pointer to it will be used in the method.
	  */
	(&emp1).changeAge(32)
	// emp1.changeAge(44) is valid too
	 fmt.Println("\n", emp1)

}