package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

func index(res http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(res, "Welcome")
}

func main() {

	indexHandler := http.HandlerFunc(index)
	logFile, err := os.OpenFile("server.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err!=nil{
		panic(err)
	}
	defer logFile.Close()

	http.Handle("/", handlers.LoggingHandler(logFile, handlers.CompressHandler(indexHandler)))

	server := &http.Server{
		Addr: ":8080",
	}
	server.ListenAndServe()
}
