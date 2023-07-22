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
	// create wait group
	var wg sync.WaitGroup

	// wait for one groutime
	wg.Add(1)
   
	/**
	Wrap call to an anonymous. This syntax creates a function 
	and immediately invokes it
	*/
	go func(){
		count("dogs")
		wg.Done()
	}()

	go func(){
		count("cats")
		wg.Done()
	}()

    // call wait at the end of the main function
	wg.Wait()
}

func count(something string){
	for i := 1; i <=10; i++ {
		fmt.Println(i, something)

		time.Sleep(time.Millisecond * 500) // sleep for half a second
	}
}