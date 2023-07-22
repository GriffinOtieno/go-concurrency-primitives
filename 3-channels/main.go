package main

import (
	"fmt"
	"time"
)

func main(){
	channel := make(chan string)
	go count("dogs", channel)

	// receive the message
	message := <-channel
	fmt.Println(message)
}

func count(something string, c chan string ){
	for i := 1; i <=10; i++ {

		// place a message on the channel
		c <- something
         
		// fmt.Println(i, something)
		time.Sleep(time.Millisecond * 500)
	}
}