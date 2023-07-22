package main

import (
	"fmt"
	"sync"
	"time"
)

func main(){
	var wg sync.WaitGroup

	wg.Add(1)
   
	go func(){
		count("dogs")
		wg.Done()
	}()

	go func(){
		count("cats")
		wg.Done()
	}()


	wg.Wait()
}

func count(something string){
	for i := 1; i <=10; i++ {
		fmt.Println(i, something)
		time.Sleep(time.Millisecond * 500)
	}
}