// minimal echo and counter server

package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

// the main function connects any URLs with a path beginning with "/" to a handler and starts a server which is listening for requests on port 8000
func main() {
	// when a request arrives, its given to the handeler
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	// a Request is a struct of type http. the URL of the incoming request is one of those fields
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
	// we add these mu.Lock() and mu.Unlock() so that we do not face a race condition and that only one goroutine accesses a variable at a time
}