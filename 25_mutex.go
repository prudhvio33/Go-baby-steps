package main

import (
	"sync"
	"fmt"
)

/*

	A Mutex is used to provide a locking mechanism to ensure that only one Goroutine is running the critical section of
	code at any point of time to prevent race condition from happening.

	Mutex is available in the sync package. There are two methods defined on Mutex namely Lock and Unlock. Any code that
	is present between a call to Lock and Unlock will be executed by only one Goroutine, thus avoiding race condition.

	mutex.Lock()
	x = x + 1
	mutex.Unlock()

	If one Goroutine already holds the lock and if a new Goroutine is trying to acquire a lock, the new Goroutine will be
	blocked until the mutex is unlocked.
*/

var x = 0

func increment(wg *sync.WaitGroup, mutex *sync.Mutex) {
	mutex.Lock()
	x = x + 1
	mutex.Unlock()
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	var mutex sync.Mutex

	for i:=0; i < 1000; i ++ {
		wg.Add(1)
		go increment(&wg, &mutex)
	}
	wg.Wait()
	fmt.Println(x)
}