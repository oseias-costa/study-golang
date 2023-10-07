package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloWord(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hellow Word\n")
}

func main() {
	http.HandleFunc("/", helloWord)
	fmt.Println("Server started and listening on localhost:9003")
	log.Fatal(http.ListenAndServe(":9003", nil))
}
