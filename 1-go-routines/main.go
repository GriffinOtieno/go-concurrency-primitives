package main

import (
	"fmt"
	"time"
)

/**
  Concurrency is a big selling point in go
  Its not the same thing as paralelism

  To run things in parallel means running things
  at exactly the same time on a multicore processor

   Concurrency is about breaking up a program into
   independently executing tasks. Can be parallelized
*/

func main(){
	// create 2 go routines
	go count("dogs")
	go count("cat")

	/**
		In go when the main go routine finishes,
		the program exits regardless of what the other
		go routines might be doing

		You can stop the main function from immediately terminating
	*/
    fmt.Scanln() // its a blocking call, will wait for user input
}

func count(something string){
	for i := 1; true; i++ { // infinite for loop
		fmt.Println(i, something)
		time.Sleep(time.Millisecond * 500)
	}
}