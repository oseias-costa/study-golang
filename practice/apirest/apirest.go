package apirest

import (
	"fmt"
	"log"
	"net/http"
)

func helloWord(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Word")
}

func Run(addr string) {
	http.HandleFunc("/", helloWord)
	fmt.Println("Server listen on 9003")
	log.Fatal(http.ListenAndServe(addr, nil))
}
