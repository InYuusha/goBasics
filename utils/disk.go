package main

import (
	"fmt"
	"log"

	"github.com/shirou/gopsutil/disk"
)

func getDisks() *disk.UsageStat {
	tmp, err := disk.Usage("/")
	if err != nil {
		log.Printf("%v", err)
	}
	return tmp

}
func main() {

	v := getDisks()

	fmt.Println(v)

}
