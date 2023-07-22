package main

import (
	"fmt"
	"time"
)
// A goroutine is simply a lightweight thread

func main(){
	go count("dogs")
	go count("cat")

    fmt.Scanln()

}

func count(something string){
	for i := 1; true; i++ {
		fmt.Println(i, something)
		time.Sleep(time.Millisecond * 500)
	}
}