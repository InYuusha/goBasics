package main

import (
	"fmt"
	"log"
	"time"

	"github.com/shirou/gopsutil/mem"
)

func getPerc() *mem.VirtualMemoryStat {
	memory, err := mem.VirtualMemory()
	if err != nil {
		log.Printf("%v", err)
	}
	return memory
}

func main() {
	for {
		time.Sleep(time.Second*2)
		v := getPerc()

		perc:=v.UsedPercent

		fmt.Println(int(perc))
	}
}
