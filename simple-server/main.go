package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/hello", hello)

	fmt.Println("Now running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
