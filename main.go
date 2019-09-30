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

func getenv(key, fallback string) string {
  if value, ok := os.LookupEnv(key); ok {
    return value
  }
  return fallback
}

func main() {
  port := ":" + getenv("HELLO_PORT", "8080")
  fmt.Printf("go-hello running in port %s\n", port)
  http.HandleFunc("/", hello)
  log.Fatal(http.ListenAndServe(port, nil))
}
