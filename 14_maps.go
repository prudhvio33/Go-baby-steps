package main

import "fmt"

/*
	A map is a builtin type in Go which associates a value to a key. The value can be retrieved using the corresponding key.

	personSalary := make(map[string]int)  It creates string keys and int values

	Similar to slices, maps are reference types. When a map is assigned to a new variable, they both point to the same
	internal data structure. Hence changes made in one will reflect in the othe
 */

func main() {
	var personSalary map[string]int

	/*
		personSalary := map[string]int {
        "steve": 12000,
        "jamie": 15000,
    }
	 */
	if personSalary == nil {
		fmt.Println("map is nil. Going to make one.")
		personSalary = make(map[string]int)
	}

	personSalary["krishna"] = 60000
	fmt.Println(personSalary)

	fmt.Println(personSalary["krishna"]) // If elem is not present it will return 0

	//CHECK IF KEY IS IN A MAP
	value, ok := personSalary["krishna"]

	fmt.Println(value, ok)  // value in the key and true/false
	value1, ok1 := personSalary["chaitanya"]
	fmt.Println(value1, ok1)

	for keys, values := range personSalary{
		fmt.Println(keys, values)
	}
}
