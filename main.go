package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("hello, world!")
		fmt.Println("CF_INSTANCE_IP", os.Getenv("CF_INSTANCE_IP"))
		fmt.Fprintf(w, "CF_INSTANCE_IP", os.Getenv("CF_INSTANCE_IP"))
	})
	http.HandleFunc("/panic", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Content-Length", "100000")
		panic("I panicked")
	})


	log.Fatal(http.ListenAndServe(":8080", nil))
}
