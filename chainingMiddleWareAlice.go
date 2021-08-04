package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/justinas/alice"
)

func firstMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		log.Printf("Iam First Middleware")
		next.ServeHTTP(res, req)
	})
}
func secondMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		log.Printf("Iam Second Middleware")
		next.ServeHTTP(res, req)
	})
}
func index(res http.ResponseWriter, req *http.Request) {
	log.Printf("Iam Index")
	fmt.Fprintf(res, "Hello Iam Index Route")

}

func main() {
	
	//Normal
	//http.Handle("/", firstMiddleware(secondMiddleware(http.HandlerFunc(index))))

	//Alice
	http.Handle("/", alice.New(firstMiddleware, secondMiddleware).ThenFunc(http.HandlerFunc(index)))
	server := &http.Server{
		Addr: ":8080",
	}
	server.ListenAndServe()
}
