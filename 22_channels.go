package main

import "fmt"

/*
	Channels can be thought as pipes using which Goroutines communicate. Similar to how water flows from one end to another
	in a pipe, data can be sent from one end and received from the another end using channels.

	Each channel has a type associated with it. This type is the type of data that the channel is allowed to transport.
	No other type is allowed to be transported using the channel.

	chan T is a channel of type T

	The zero value of a channel is nil. nil channels are not of any use and hence the channel has to be defined using
	make similar to maps and slices.
*/
func hello(done chan bool) {
	fmt.Println("Hello from goroutine")
	done <- true
}

func producer(channel chan int) {
	for i:=0; i<10; i++ {
		channel <- i
	}

	close(channel)
}

func main()  {
	var done chan bool
	if done == nil {
		done = make(chan bool)
	}

	// data := <- a // read from channel a

	// a <- data // write to channel a

	/*
	Sends and receives to a channel are blocking by default. What does this mean? When a data is sent to a channel, the
	control is blocked in the send statement until some other Goroutine reads from that channel. Similarly when data is
	read from a channel, the read is blocked until some Goroutine writes data to that channel.

	This property of channels is what helps Goroutines communicate effectively without the use of explicit locks or
	conditional variables that are quite common in other programming languages.
	 */

	 go hello(done)
	 <- done	// This will be blocked until something is written to this channel
	 fmt.Println("Main")

	 /*
	 	DEADLOCKS
	 	---------
	 	One important factor to consider while using channels is deadlock. If a Goroutine is sending data on a channel,
	 	then it is expected that some other Goroutine should be receiving the data. If this does not happen, then the
	 	program will panic at runtime with Deadlock.

		Similarly if a Goroutine is waiting to receive data from a channel, then some other Goroutine is expected to write
	 	data on that channel, else the program will panic.
	  */

	//ch := make(chan int)
	//ch <- 5

	/*
		CLOSING CHANNEL
		---------------
		Senders have the ability to close the channel to notify receivers that no more data will be sent on the channel.

		Receivers can use an additional variable while receiving data from the channel to check whether the channel has been closed.

		v, ok := <- ch
		ok is true, then channel is alive
		ok is false, channel is dead
	 */
	 channel := make(chan int)
	 go producer(channel)

	 for {
	 	v, ok := <- channel
	 	if ok == false {
	 		break
		}
		fmt.Println(v)
	 }

	for v := range channel {
		fmt.Println("Received ",v)
	}
}