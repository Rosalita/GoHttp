// This is a minimal "echo" and counter webserver
// while server running, navigating to localhost:8000/any_path, will echo the url path and increment a counter
// wnavigating to localhost:8000/count, will display the counter
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	// This server has two handlers, the request url determines which one is called
	http.HandleFunc("/", handler)      // connect a handler function to incoming urls that start with "/"
	http.HandleFunc("/count", counter) // connect a counter function to incoming urls that start with "/count"
	// start a server listening for incomming requests on port 8000
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the requested url and increments a counter.
func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock() //to prevent two requests updating at the same time, lock before incrementing
	count++
	mu.Unlock() // and unlock after incrementing, this prevents "race condition" bugs

	// extract the path component from the request url and send it back as the response
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// counter function displays the counter
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}
