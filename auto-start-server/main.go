package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"runtime"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}

func open(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/hello", hello)

	url := "http://localhost:8080"
	go open(url)

	fmt.Println("Now running on", url)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
