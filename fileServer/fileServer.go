package main

import (
	"log"
	"net/http"
)

func main() { 
	//build the file
	fs := http.FileServer(http.Dir("server/public"))
	log.Fatal(http.ListenAndServe(":8080", fs))
}
