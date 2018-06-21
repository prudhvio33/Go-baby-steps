package employee

import "fmt"

/*
	Names starting with capitals can be imported by other programs, to make them private name them with small chars
 */

type Employee struct {
	Name string
	Age  int
	Salary int
}

func (e Employee) ShowSalary() {
	fmt.Printf("Salary %d", e.Salary)
}

type employeeV2 struct {
	firstName   string
	lastName    string
	totalLeaves int
	leavesTaken int
}

func New(firstName string, lastName string, totalLeave int, leavesTaken int) employeeV2 {
	e := employeeV2 {firstName, lastName, totalLeave, leavesTaken}
	return e
}

func (e employeeV2) LeavesRemaining() {
	fmt.Printf("%s %s has %d leaves remaining", e.firstName, e.lastName, (e.totalLeaves - e.leavesTaken))
}