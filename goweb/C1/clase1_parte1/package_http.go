package main

import (
	"fmt"
	"net/http"
)

func holaHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "pong\n")
}

func main() {
	http.HandleFunc("/ping", holaHandler)
	http.HandleFunc("/pang", holaHandler)
	http.ListenAndServe(":8080", nil)
}
