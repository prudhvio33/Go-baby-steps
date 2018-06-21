package main

import "fmt"

/*
	defer keyword is used to execute a function call just before the function it is present in returns

	Multiple defer calls will be executed in LIFO order

	Defer is used in places where a function call should be executed irrespective of the code flow.
*/

func finished()  {
	fmt.Println("Finished finding large numbers")
}

func findLargestNumber(nums []int) {
	defer finished()
	max := nums[0]
	for _, v := range nums {

		if v > max {
			max = v
		}

	}

	fmt.Println("Max number is ", max)
}

func printA(a int) {
	fmt.Println("value of a in deferred function", a)
}

func main() {
	nums := []int{1,2,4,100,3}
	findLargestNumber(nums)

	/*
		Max number is  100
		Finished finding large numbers

	 */

	a := 5
	defer printA(a)
	a = 10
	fmt.Println("value of a before deferred function call", a)

	/*
		value of a before deferred function call 10
		value of a in deferred function 5
	 */
}