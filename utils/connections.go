package main

import (
	"fmt"
	"log"
	"time"

	"github.com/shirou/gopsutil/net"
)

func getConn() []net.ConnectionStat {
	conn, err := net.Connections("inet")
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
