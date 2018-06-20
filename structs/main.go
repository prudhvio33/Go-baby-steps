package structs

import "fmt"
import "./computers"

func main()  {
	var spec computers.Spec
	spec.Maker = "apple"
	spec.Price = 50000
	fmt.Println("Spec:", spec)
}