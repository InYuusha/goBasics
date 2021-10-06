package main

import (
	"fmt"
	"log"
	"time"

	"github.com/shirou/gopsutil/docker"
)

func getInfo()interface{}{
	list, err := docker.GetDockerIDList()
	if err != nil {
		log.Println(err)
	}
	return list
}
func main() {
	for {
		time.Sleep(time.Second * 2)
		v := getInfo()

		fmt.Println(v)
	}
}
