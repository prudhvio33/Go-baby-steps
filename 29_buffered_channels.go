package main

/*
	Go also has buffered channels that can store more than a single element. For example, ch := make(chan bool, 4), here
	we create a channel that can store 4 boolean elements. So in this channel, we are able to send 4 elements into it
	without blocking, but the goroutine will be blocked when you try to send a fifth element and no goroutine receives it.
*/

import "fmt"

func main() {
	c := make(chan int, 2) // change 2 to 1 will have runtime error, but 3 is fine
	c <- 1
	c <- 2
	fmt.Println(<-c)
	fmt.Println(<-c)
}