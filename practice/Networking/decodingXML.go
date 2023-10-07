package main

import (
	"encoding/xml"
	"fmt"
)

type person struct {
	XMLName	   xml.Name `xml:"persondata"` 
	Name       string   `json:"firstname"`
	Address    string   `json:"addr"`
	Age        int      `json:"age,attr"`
	FaveColors []string `json:"favecolors"`
}

func main(){
	//create same data
	xmldata := `
		<persondata age="29">
			<firstname>John</firstname>
			<lastName>Public</lastName>
			<addr>Otavio rocha</addr>
			<favecolors>white</favecolors>
			<favecolors>red</favecolors>
		</persondata>
	`

	//data will be decoded into a person sctruct
	var p person

	// use the Unmarshal function to decode raw xml data
	xml.Unmarshal([]byte(xmldata), &p)
	fmt.Printf("%#v\n", p)
}