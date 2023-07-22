package main

import (
	"fmt"
	"sync"
	"time"
)

/**
Wait groups are used to wait for
go routines to complete execution
*/

func main(){
	var wg sync.WaitGroup // create wait group

	// indicate there is one go routine to wait for
	// just a counter - increment the counter (how many go routines are running)
	wg.Add(1)
   
	// create an anonymous function and immediately invoke it
	go func(){
		count("dogs")
		wg.Done() //decrement the counter by when the go routine finishes
	}()

	// create an anonymous function and immediately invoke it
	go func(){
		count("cats")
		wg.Done() 
	}()

	// call wait at the end of the main function. Block until the count is 0
	wg.Wait() // will wait if the go routines havent finished
}

func count(something string){
	for i := 1; i <=10; i++ {
		fmt.Println(i, something)
		time.Sleep(time.Millisecond * 500) // sleep for half a second
	}
}