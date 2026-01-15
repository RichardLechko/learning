// minimal echo server

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

var keyVals = make(map[string]string)
var Port = 8000
var localURL = "localhost:"

func main() {
	input := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter a Port Number: ")
	input.Scan()
	line := input.Text()
	i, err := strconv.Atoi(line)
	if err != nil {
		fmt.Println("Cannot convert input to an Integer. Please try again.")
		os.Exit(1)
	}
	Port = i
	localURL += strconv.Itoa(Port)
	fmt.Printf("Server running on: %s\n", localURL)
	http.HandleFunc("/", handler)
	http.HandleFunc("/keys/", customHandler)
	http.HandleFunc("/keys", listAllKeys)
	log.Fatal(http.ListenAndServe(localURL, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)

	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	fmt.Fprintf(w, "Body = %q\n", r.Body)

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

func customHandler(w http.ResponseWriter, r *http.Request) {

	key := r.URL.Path[6:]

	switch r.Method {
	case "PUT":
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "could not ready body", http.StatusBadRequest)
			return
		}

		var data struct{ Value string }
		if err := json.Unmarshal(body, &data); err != nil {
			log.Fatalf("JSON unmrashaling failed: %s", err)
		}
		keyVals[key] = data.Value

		log.Printf("Received: %s%T\tKey = %s\n", body, body, key)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "OK\n")
	case "GET":
		value, ok := keyVals[key]
		if !ok {
			http.Error(w, "Key does not exist", http.StatusNotFound)
			break
		}
		log.Printf("Received GET Request hitting %v/%v\n", localURL, key)

		fmt.Fprintf(w, "%s\n", value)
	case "DELETE":
		_, ok := keyVals[key]
		if !ok {
			http.Error(w, "Key does not exist", http.StatusNotFound)
			break
		}
		delete(keyVals, key)
		log.Printf("Deleted %v", key)
		fmt.Fprintf(w, "Deleted\n")
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

}

func listAllKeys(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		for key, value := range keyVals {
			fmt.Fprintf(w, "%s\t%s\n", key, value)
		}
		log.Printf("Received GET Request to list all keys.")
	default:
		fmt.Printf("Method Used: %q\n", r.Method)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)

	}
}
