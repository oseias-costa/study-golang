package main

import "net/http"

func main(){
	http.HandleFunc("/person/", func(w http.ResponseWriter, r *http.Request){
		http.Error(w, "Not implemented", http.StatusInternalServerError)
	})

	http.ListenAndServe(":8080", nil)

}