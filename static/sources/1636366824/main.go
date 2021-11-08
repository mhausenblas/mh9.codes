package main

import (
	"fmt"
	"log"
	"net/http"
)

func handlePing(w http.ResponseWriter, r *http.Request) {
	input := r.URL.Query().Get("this")
	fmt.Fprintf(w, input)
}

func main() {
	http.HandleFunc("/echo", handlePing)
	log.Fatal(http.ListenAndServe(":4242", nil))
}
