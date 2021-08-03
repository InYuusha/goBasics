//unbuffered channel (capacity not defined)
package main

import (
	"fmt"
	"time"
)

func compute(c chan string) {
	fmt.Println("Executing Go Routine")
	time.Sleep(time.Second)
	c <- "Hello World"
	fmt.Println("Finished Executing")
}

func main() {
	c:=make(chan string)
	defer close(c)

	go compute(c) //only one goroutine will finish
	go compute(c)	//since we are recieving only one output
					//other go routine will be waiting for its sent value to be recieved
					//and will be terminated
	k := <-c
	fmt.Println(k)

}
