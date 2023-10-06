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

func decodeExample(){
	data := []byte(`
		{
			"fullname": "John Q Public".
			"addr": "987 Main St",
			"age": 45,
			"favecolors": ["Purple", "Write", "Gold"]
		}
	`)

	var p person

	//test to see if the Json is valid, and then decodeExample
	valid := json.Valid(data)
	if valid {
		json.Unmarshal(data, &p)
		fmt.Printf("%#v\n", p)
	}

	//data can also be decoded into a map structure
	var m map[string]interface{}

	//Unmarshal into a map
	json.Unmarshal(data, &m)
	fmt.Printf("%#v\n", m)

	for k, v := range m {
		fmt.Printf("key (%v), value (%T : %v)\n", k, v, v)
	}
}

