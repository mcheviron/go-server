package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() error: %v\n", err)
		return
	}
	fmt.Fprintf(w, "POST Request sucsseful\n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name: %v\n", name)
	fmt.Fprintf(w, "Address: %v\n", address)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello!\n")
}

func main() {
	// Will go fetch the index.html from the static folder because Go webservers are trained to look for it
	// ./static will serve as the root of our server
	webserver := http.FileServer(http.Dir("./static"))

	// This will handle requests to the root server by just passing it to the webserver, which is trained to look
	// for the index.html, thus it will simply return the contents of the index.html
	http.Handle("/", webserver)

	// The next two functions map specific handlers/handle functions to specific requests
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Starting on port 8080")

	// Will listen on TCP port 8080
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
