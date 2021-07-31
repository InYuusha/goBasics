package main

import (
	"fmt"
	"time"
)

func halt(n int){
	for i:=0;i<n;i++{
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

func main(){
	go halt(5)
	go halt(5)

	fmt.Scanln() //wait
}