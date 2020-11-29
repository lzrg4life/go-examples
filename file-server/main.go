// Taken from https://gist.github.com/paulmach/7271283
/*
Serve is a very simple static file server in go
Usage:
	-p="8100": port to serve on
	-d=".":    the directory of static files to host
Navigating to http://localhost:8100 will display the index.html or directory
listing file.
*/
package main

import (
	"flag"
	"log"
	"net/http"
	"os/exec"
)

func open(url string) {
	exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
}

func main() {
	port := flag.String("p", "8100", "port to serve on")
	directory := flag.String("d", ".", "the directory of static file to host")
	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir(*directory)))

	strport := ":" + *port
	url := "http://localhost" + strport
	open(url)

	log.Printf("Serving %s on %s\n", *directory, url)
	log.Fatal(http.ListenAndServe(strport, nil))
}
