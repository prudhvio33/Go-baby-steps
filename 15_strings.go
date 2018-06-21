package main

import (
	"fmt"
	"unicode/utf8"
)

/*
	A string in Go is a slice of bytes.



 */
func printBytes(s string) {
	for i:= 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i]) //68 65 6c 6c 6f 20 77 6f 72 6c 64
	}
}

func printChars(s string) {
	fmt.Println("\n")
	for index, value := range s {
		fmt.Printf("%c starts at byte %d\n", value, index)
	}
}

func mutate(s []rune) string {
	s[0] = 'a'
	return string(s)
}

func main() {
	printBytes("hello world")
	printChars("hello world")

	// LENGTH
	fmt.Println(len("hello"))
	fmt.Println(utf8.RuneCountInString("hello"))

	//Strings are immutable in Go. Once a string is created its not possible to change it.
	//To workaround this string immutability, strings are converted to a slice of runes. Then that slice is mutated with
	// whatever changes needed and converted back to a new string.

	h := "hello"
	fmt.Println(mutate([]rune(h)))

	// or
	s := "hello"
	c := []byte(s)  // convert string to []byte type
	c[0] = 'c'
	s2 := string(c)  // convert back to string type
	fmt.Printf("%s\n", s2)
}
