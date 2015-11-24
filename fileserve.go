// Start a HTTP server serving the files in a directory.

// Usage:

// $ go run fileserve.go -d <directory> -p <port>

// Examples:

// By default, the files in the current directory will be served on port 8080:
// $ go run fileserve.go 

// To specify a different directory and port:
// $ go run fileserve.go -d=/tmp -p=9090

package main

import (
	"net/http"
	"flag"
	"log"
)

type myType struct {}

func (t *myType) ServeHTTP(w http.ResponseWriter, r *http.Request){
	log.Printf("Request recieved")
	http.FileServer(http.Dir("/tmp"))
}

func main() {
	port := flag.String("p", "8080", "Port to serve on")
	//directoryToServe := flag.String("d", ".", "Directory to serve")
	flag.Parse()	
	listenOn := "127.0.0.1:" + *port
	// ListenAndServe is blocking and runs each handler in a goroutine, so can serve multiple
	// requests simultaneously
	// http.FileServer: https://golang.org/pkg/net/http/#FileServer
	myTypeh := &myType{}
	http.Handle("/", myTypeh)
	err := http.ListenAndServe(listenOn, nil) //http.FileServer(http.Dir(*directoryToServe)));
	if (err != nil) {
		log.Fatal("ListenAndServe", err)
	}
}
