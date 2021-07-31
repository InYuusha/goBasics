package main

import "fmt"

func sendData(s chan int){
	s<-5  //push in channel
	fmt.Scanln() //wait here (thread pause)
	fmt.Print("Finished") //wont reach here
}
func main(){
	
	s:=make(chan int)
	go sendData(s)
	k:=<-s //recieve  (other end of channel)
	fmt.Print(k)


}