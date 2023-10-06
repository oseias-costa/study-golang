package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name       string   `json:"fullname"`
	Address    string   `json:"addr"`
	Age        int      `json:"age"`
	FaveColors []string `json:"favecolors"`
}

func encodeExample() {
	// create some people data
	people := []person{
		{"Jane Doe", "123 anywgere street", 35, nil},
		{"John Public", "45 Were bkvs", 31, []string{"purple", "Yellow"}},
	}

	result, err := json.MarshalIndent(people, "", "\t")
	if err != nil {
		return
	}
	fmt.Printf("%s\n", result)
}

func main() {
	encodeExample()
}
