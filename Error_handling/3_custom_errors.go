package main

import (
	"errors"
	"math"
	"fmt"
)

/*
	Simplest way of creating custom errors is by using New function in errors package
 */

func circleArea(radius float64) (float64, error) {
	if radius < 0 {
		return 0, errors.New("Area calculation failed, radius is less than zero")
	}
	return math.Pi * radius * radius, nil
}

func main() {
	radius := -20.0
	area, err := circleArea(radius)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Area of circle %0.2f", area)

	// Customize error function
	fmt.Errorf("Area calculation failed, radius %0.2f is less than zero", radius)
}