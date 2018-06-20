package main

import (
	"time"
	"fmt"
)

/*
	The select statement is used to choose from multiple send/receive channel operations. The select statement blocks
	until one of the send/receive operation is ready. If multiple operations are ready, one of them is chosen at random.

	Lets assume we have a mission critical application and we need to return the output to the user as quickly as possible.
	The database for this application is replicated and stored in different servers across the world. Assume that the
	functions server1 and server2 are in fact communicating with 2 such servers. The response time of each server is
	dependant on the load on each and the network delay. We send the request to both the servers and then wait on the
	corresponding channels for the response using the select statement. The server which responds first is chosen by the
	select and the other response is ignored. This way we can send the same request to multiple servers and return the
	quickest response to the user.

*/

func server1(ch chan string) {
	time.Sleep(6 * time.Second)
	ch <- "from server1"
}
func server2(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "from server2"

}
func main() {
	output1 := make(chan string)
	output2 := make(chan string)
	go server1(output1)
	go server2(output2)
	select {
		case s1 := <-output1:
			fmt.Println(s1)
		case s2 := <-output2:
			fmt.Println(s2)
		default:
			fmt.Println("no value received")

	}
}