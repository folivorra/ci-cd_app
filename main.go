package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, test of ci/cd on github action is working!")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server starts on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
