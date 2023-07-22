package main

import "fmt"

func main(){
	channel := make(chan string, 1)
	
	channel <- "First Message"
	message := <- channel
	fmt.Println(message)

    channel <- "Second Message"
	message = <- channel
	fmt.Println(message)


}
