package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", handleHello)
	http.ListenAndServe(":8080", nil)
}

func handleHello(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, "<h1>Hello World, ආයුබෝවන්</h1>")
}
