package main

import (
    "fmt"
    "net/http"
)

func greet(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, you've reached my microservice!")
}

func main() {
    http.HandleFunc("/", greet)
    fmt.Println("Server listening on port 8080")
    http.ListenAndServe(":8080", nil)
}
