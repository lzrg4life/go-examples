package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}

func open(url string) {
	exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/hello", hello)

	url := "http://localhost:8080"
	open(url)

	fmt.Println("Now running on", url)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
