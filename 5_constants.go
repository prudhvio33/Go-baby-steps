package main

import "math"

func main()  {
	const a = 10

	var b = math.Sqrt(10)

	/* The value of a constant should be known at compile time. Hence it cannot be assigned to a value returned by a
	function call since the function call takes place at run time
	*/
	const c = math.Sqrt(10)


}
