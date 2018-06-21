package main

import (
	"./employee"
	"fmt"
)

func main()  {
	e := employee.Employee{"krishna Chaitanya", 24, 10000000}
	fmt.Println(e)
	e.ShowSalary()

	e2 := employee.New("Sam", "Adolf", 30, 20)
	e2.LeavesRemaining()
}
