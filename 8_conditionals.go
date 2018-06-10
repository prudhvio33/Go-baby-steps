package main

import "fmt"

/*
	if condition {
	} else if condition {
	} else {
	}

	else or else if should be in the same line as } as go lang automatically adds ; after statement and if else is in
	different line if condition may be terminated in that line
 */

func main()  {
	a := 10
	b := 12

	if a > b {
		fmt.Println(" a is greater")
	} else {
		fmt.Println("b is greater")
	}
}

