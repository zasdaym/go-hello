package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World. Path: %s\n", r.URL.Path)
	log.Println(r.URL.Path)
}

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}

func main() {
	port := ":" + getEnv("PORT", "8080")
	fmt.Printf("running on %s\n", port)
	http.HandleFunc("/", hello)
	log.Fatal(http.ListenAndServe(port, nil))
}
