package main

import (
	"fmt"
	"time"
)
/*
	Goroutines are functions or methods that run concurrently with other functions or methods. Goroutines can be thought
	of as light weight threads. The cost of creating a Goroutine is tiny when compared to a thread. Hence its common for
	Go applications to have thousands of Goroutines running concurrently.

	Advantages of Goroutines over threads
	-------------------------------------
	1) Goroutines are extremely cheap when compared to threads. They are only a few kb in stack size and the stack can
  	   grow and shrink according to needs of the application whereas in the case of threads the stack size has to be specified
	   and is fixed.

	2) The Goroutines are multiplexed to fewer number of OS threads. There might be only one thread in a program with
	  thousands of Goroutines. If any Goroutine in that thread blocks say waiting for user input, then another OS thread
		is created and the remaining Goroutines are moved to the new OS thread. All these are taken care by the runtime
		and we as programmers are abstracted from these intricate details and are given a clean API to work with concurrency.

	3) Goroutines communicate using channels. Channels by design prevent race conditions from happening when accessing
	   shared memory using Goroutines. Channels can be thought of as a pipe using which Goroutines communicate. We will
	   discuss channels in detail in the next tutorial.
*/

func numbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}
func alphabets() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}

func main()  {
	go numbers() //starts a new Goroutine. Now the hello() function will run concurrently along with the main() function
	go alphabets()

	// output: 1 a 2 3 b 4 c 5 d e

	// When one sleeps other executes
	/*
		When a new Goroutine is started, the goroutine call returns immediately. Unlike functions, the control does not
		wait for the Goroutine to finish executing. The control returns immediately to the next line of code after the
		Goroutine call and any return values from the Goroutine are ignored.

		The main Goroutine should be running for any other Goroutines to run. If the main Goroutine terminates then the
		program will be terminated and no other Goroutine will run.
	 */
	time.Sleep(1 * time.Second) // This is just a hack, we will be using channels to block the main goroutine
	fmt.Println("Main function")


}

