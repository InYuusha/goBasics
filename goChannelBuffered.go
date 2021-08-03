package main

import (
	"fmt"
	"time"
)

func compute(c chan string) {
	fmt.Println("Starting Execution")
	time.Sleep(time.Second)
	c<-"Hello world"
	fmt.Println("Finished Execution")

}

func main() {
	c := make(chan string, 2)
	defer close(c)

	go compute(c)
	go compute(c)

	k:=<-c
	fmt.Println(k)

	time.Sleep(time.Second)

}
