package main

import "fmt"

/*

	AA TYPES IMPLEMENT EMPTY INTERFACE interface{}
	The general definition of an interface in the Object oriented world is "interface defines the behaviour of an object".
	It only specifies what the object is supposed to do. The way of achieving this behaviour (implementation detail) is upto the object.

	Interface specifies what methods a type should have and the type decides how to implement these methods.

	For example WashingMachine can be a interface with method signatures Cleaning() and Drying(). Any type which provides
	definition for Cleaning() and Drying() is said to implement the WashingMachine interface.
*/

// Interface declaration
type VowelsFinder interface {
	FindVowels() []rune
}

type MyString string

/*

	we add the method FindVowels() []rune to the receiver type MyString. Now MyString is said to implement the interface
	VowelsFinder. This is quite different from other languages like Java where a class has to explicitly state that it
	implements a interface using the implements keyword. This is not needed in go and go interfaces are implemented implicitly
	if a type contains all the methods declared in the interface.
*/

////MyString implements VowelsFinder
func (ms MyString)FindVowels() []rune {
	var vowels []rune
	for _, rune := range ms {
		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
			vowels = append(vowels, rune)
		}
	}
	return vowels

}

func main(){
	name := MyString("Krishna Chaitanya")
	var v VowelsFinder
	v = name // possible since MyString implements VowelsFinder
	fmt.Printf("Vowels are %c", v.FindVowels())
	fmt.Printf("Vowels are %c", name.FindVowels())

}
