package main

import (
	"fmt"
	"time"
)

/**
  Concurrency is a big selling point in go
  Its not the same thing as paralelism
*/

func main(){

	// create 2 go routines
	go count("dogs")
	go count("cat")

	// stop the main function from immediately terminating
    fmt.Scanln()

}

func count(something string){
	for i := 1; true; i++ {
		fmt.Println(i, something)
		time.Sleep(time.Millisecond * 500)
	}
}