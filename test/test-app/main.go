package main

import (
  "fmt"
  "net/http"
)

const (
  port = ":8080"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "Hello, world !")
}

func main() {
  fmt.Printf("Started server on port %v.\n", port)
  http.HandleFunc("/", HelloWorld)
  http.ListenAndServe(port, nil)
}
