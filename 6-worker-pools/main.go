package main

import "fmt"

/**
  A pattern where you have a queue
  of work to be done and multiple concurrent
  workers pulling items of the queue
*/

func main(){
   jobs := make(chan int, 100)
   results := make(chan int, 100)

   // create a worker as a concurrent go routine
   // this is still inefficient, as it maxes the cpu to 99%
   // trying to calculate fib numbers
   go worker(jobs, results)

   // you can add more workers
   go worker(jobs, results)
   go worker(jobs, results)
   go worker(jobs, results)
   // uses almost 400% cpu

   //fill out the jobs channel with 100 numbers
   //since its buffered, we are not going to block
   for i:= 0; i <100; i ++ {
	 jobs <- i
   }

   close(jobs)

   // output results to terminal
  
   for j := 0; j< 100; j++ {
	 fmt.Println(<-results)
   }
}

/**
	Create a worker that takes two channels
	one channel of jobs to do, and another to send 
	results on

	Instead of creating bidirectional channels,
	specify that we will only ever receive from the jobs
	channel and receive from the result channel

	This reduces the chance of bugs
*/

func worker(jobs <-chan int, results chan<- int){
   // the worker will pull one job at a time and push to results channel
	//consume items from this queue
	for n := range jobs {
		// calculate the nth fib number and send it
		// on the results channel
		results <- fib(n)
    }
}

/**
fib algorithm to calculate the nth
fibonnacci number and return it
*/

func fib(n int) int {
	// if n is 0 or 1
	if n <= 1 {
		return n
	}

	// return  the sum of the previous 2 fib numbers
	return fib(n-1) + fib(n-2)
}
