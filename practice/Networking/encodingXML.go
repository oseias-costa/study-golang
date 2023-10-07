package main

import (
	"encoding/xml"
	"fmt"
)

type person struct {
	XMLName	   xml.Name  `xml:"persondata"` 
	Name       string   `json:"fullname"`
	Address    string   `json:"addr"`
	Age        int      `json:"age"`
	FaveColors []string `json:"favecolors"`
}

func main(){
	//declare some sample data
	peaple := []person{
		{FirstName: "Jane", LastName: "Doe", Address: "455 Something", FaveColors: []string{"Blue", red}},
		{FirstName: "Oséias", LastName: "Costa", Address: "Otavio rocha", FaveColors: []string{"white", red}},
	}

	//the MarchalIndent function indents the XML text
	result, err := xml.MarchalIndent(peaple, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s%s\n", xml.Header, result)
}