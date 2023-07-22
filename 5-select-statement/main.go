package main

import (
	"fmt"
	"time"
)

func main(){

	// make 2 channels
	c1 := make(chan string)
	c2 := make(chan string)

	// goroutine to send to first channel every 1s
	go func(){
		for {
			c1 <- "Every 1s"
			time.Sleep(time.Second * 1)
		}
	}()

	// goroutine to send to second channel every 1s
	go func(){
		for {
			c2 <- "Every 2s"
			time.Sleep(time.Second * 2)
		}
	}()

	//   fmt.Println(<-c1)
	//   fmt.Println(<-c2)

	// receive from channel 1 and 2
	for {

	//receive from whichever channel is ready
	  select {
	  	case msg1 := <- c1:
		fmt.Println(msg1)
	   
	    case msg2 := <- c2:
		fmt.Println(msg2)

	  }
	}

}
