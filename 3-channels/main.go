package main

import (
	"fmt"
	"time"
)

func main(){
	// a pipe to send/receive messages between processes
	c := make(chan string)
	go count("dogs", c)

	// receive the message from the channel

	/**
		We get deadlock, because the main function is still
		waiting to receive the channel, but nothing else is
		ever going to send the message on the channel

		Go will detect this halting problem at runtime

		The receiver cannot close a channel as it doesnt know if the 
		sender is finished
	*/

	// for {
	// 	msg, open := <- c
	// 	if !open {
	// 		break
	// 	}
	// 	fmt.Println(msg)
	// }

	for msg := range c {
		fmt.Println(msg)
	}
}

func count(something string, c chan string ){
	for i := 1; i <= 10; i++ {
		c <- something //send on the channel
		// fmt.Println(i, something)
		time.Sleep(time.Millisecond * 500)
	}

	// to solve this we can close the channel
	// is it knows that its done
	close(c)
}