package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type contextKey string

//incoming reqbody
type Request struct {
	Name string `json:"name"`
}

type Response struct {
	Message string `json:"msg"`
}

type ValidationHandler struct {
	next http.Handler
}

func newValidationHandler(next http.Handler) http.Handler {
	return ValidationHandler{next: next}
}

//Decder
func (h ValidationHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var name Request
	err := json.NewDecoder(req.Body).Decode(&name)
	if err != nil {
		http.Error(res, "Bad Request", http.StatusBadRequest)
	}
	c := context.WithValue(req.Context(), contextKey("name"), name.Name)
	req = req.WithContext(c)
	h.next.ServeHTTP(res, req)

}
func newHelloWorldHandler() http.Handler {
	return helloWorld{}
}

type helloWorld struct{}

//Encoder
func (h helloWorld) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	name := req.Context().Value(contextKey("name")).(string)
	message := Response{Message: "Hello " + name}
	err := json.NewEncoder(res).Encode(message)

	if err != nil {
		http.Error(res, "Bad Request", http.StatusBadRequest)
	}

}

func main() {

	const port=8080

	handler := newValidationHandler(newHelloWorldHandler())
	http.Handle("/", handler)

	log.Printf("Server Starting on port %v",port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
