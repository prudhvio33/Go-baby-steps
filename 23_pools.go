package main

import (
	"sync"
	"fmt"
	"time"
	"math/rand"
)

/*
	WaitGroup
	---------
	A WaitGroup is used to wait for a collection of Goroutines to finish executing. The control is blocked until all
	Goroutines finish executing. Lets say we have 3 concurrently executing Goroutines spawned from the main Goroutine.
	The main Goroutines needs to wait for the 3 other Goroutines to finish before terminating. This can be accomplished
	using WaitGroup.

	WorkerPool
	----------
	In general, a worker pool is a collection of threads which are waiting for tasks to be assigned to them. Once they
	finish the task assigned, they make themselves available again for the next task.

		The following are the core functionalities of our worker pool

			Creation of a pool of Goroutines which listen on an input buffered channel waiting for jobs to be assigned
			Addition of jobs to the input buffered channel
			Writing results to an output buffered channel after job completion
			Read and print results from the output buffered channel
*/

func process(i int, wg *sync.WaitGroup) {
	fmt.Println("started Goroutine ", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d ended\n", i)
	wg.Done()
}


var jobs = make(chan Job, 10)  // Worker Goroutines listen for new tasks on the jobs buffered channel.
var results = make(chan Result, 10) // Once a task is complete, the result is written to the results buffered channel.

type Job struct {
	id       int
	randomno int
}
type Result struct {
	job         Job
	sumofdigits int
}

func digits(number int) int {
	sum := 0
	no := number
	for no != 0 {
		digit := no % 10
		sum += digit
		no /= 10
	}
	time.Sleep(2 * time.Second)
	return sum
}

func worker(wg *sync.WaitGroup) {
	for job := range jobs {
		output := Result{job, digits(job.randomno)}
		results <- output
	}
	wg.Done()
}

func createWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(results)
}

func allocate(noOfJobs int) {
	for i := 0; i < noOfJobs; i++ {
		randomno := rand.Intn(999)
		job := Job{i, randomno}
		jobs <- job
	}
	close(jobs)
}

func result(done chan bool) {
	for result := range results {
		fmt.Printf("Job id %d, input random no %d , sum of digits %d\n", result.job.id, result.job.randomno, result.sumofdigits)
	}
	done <- true
}

func main() {
	no := 3
	var wg sync.WaitGroup // is a struct type

	for i := 0; i < no; i++ {
		wg.Add(1) // When we call Add on the WaitGroup and pass it an int, the WaitGroup's counter is incremented
						// by the value passed to Add. The way to decrement the counter is by calling Done() method on
						// the WaitGroup. The Wait() methods blocks the Goroutine in which its called until the counter becomes zero.
		go process(i, &wg) // It is important to pass the address of wg. If the address is not passed,
							// then each Goroutine will have its own copy of the WaitGroup and main will not be notified
							// when they finish executing.
	}

	wg.Wait()
	fmt.Println("All go routines finished executing")


	startTime := time.Now()
	noOfJobs := 100
	go allocate(noOfJobs)
	done := make(chan bool)
	go result(done)
	noOfWorkers := 50
	createWorkerPool(noOfWorkers)
	<-done
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")
}