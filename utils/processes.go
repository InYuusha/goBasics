package main

import (
	"fmt"
	"log"
	"time"

	"github.com/shirou/gopsutil/process"
)

func getConn() interface{} {
	conn, err := process.Processes()
	if err != nil {
		log.Println(err)
	}
	return conn
}
func main() {
	for {
		time.Sleep(time.Second * 2)
		v := getConn()

		fmt.Println(v)
	}
}
