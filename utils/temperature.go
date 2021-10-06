package main

import (
	"fmt"
	"log"
	"time"

	"github.com/shirou/gopsutil/host"
)

func getTemp() interface{} {
	tmp, err := host.Info()
	if err != nil {
		log.Printf("%v", err)
	}
	return tmp

}
func main() {
	for {
		time.Sleep(time.Second * 2)
		v := getTemp()

		fmt.Println(v)
	}
}
