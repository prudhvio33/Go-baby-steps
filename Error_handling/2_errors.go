package main

import (
	"os"
	"fmt"
)

/*
	error is a type in go
 */

func main(){
	file, err := os.Open("test.txt")		// func Open(name string) (file *File, err error)
	if err != nil {
		fmt.Println(err)
		return //Important
	}
	fmt.Println(file.Name())
}
