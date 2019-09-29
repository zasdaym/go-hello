package main

import (
  "fmt"
  "log"
  "net/http"
  "os"
)

func hello(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello, World. Path: %s\n", r.URL.Path)
  log.Println("hello function executed")
}

func main() {
  os.Setenv("HELLO_PORT", "8080")
  http.HandleFunc("/", hello)
  port := ":" + os.Getenv("HELLO_PORT")
  log.Fatal(http.ListenAndServe(port, nil))
}
