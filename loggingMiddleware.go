package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func loggingHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		start := time.Now()
		log.Printf("Started %v %v", req.Method, req.URL)
		next.ServeHTTP(res, req)
		log.Printf("Completed Url %v in time %v", req.URL, time.Since(start))
	})
}
func index(res http.ResponseWriter, req *http.Request) {
	log.Println("Executing request index")
	fmt.Fprint(res, "WElcome")
}

func main() {
	indexHandler := http.HandlerFunc(index)

	http.Handle("/", loggingHandler(indexHandler))

	http.ListenAndServe(":8080", nil)

}
