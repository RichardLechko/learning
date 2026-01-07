// minimal echo server

package main

import (
	"fmt"
	"log"
	"net/http"
)

// the main function connects any URLs with a path beginning with "/" to a handler and starts a server which is listening for requests on port 8000
func main() {
	// when a request arrives, its given to the handeler
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
	// a Request is a struct of type http. the URL of the incoming request is one of those fields
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}