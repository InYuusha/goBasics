package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const port = 8080

type response struct {
	Message string `json:"msg"`
}

type request struct {
	Name string `json:"name"`
}

func hello(res http.ResponseWriter, req *http.Request) {
	var reqObj request
	err := json.NewDecoder(req.Body).Decode(&reqObj)

	if err != nil {
		http.Error(res, "Bad Request", http.StatusBadRequest)
	}
	resObj := response{Message: "Hello" + reqObj.Name}
	json.NewEncoder(res).Encode(resObj)

}

func main() {

	http.HandleFunc("/", hello)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
