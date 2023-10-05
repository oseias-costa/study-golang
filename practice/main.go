package main

import (
	"fmt"
)

const newString = "string"

func main() {
	var asString string = "isso é uma string"

	var anInteger int = 42
	var defaultInt int

	myString := "This is also a string"

	fmt.Println(asString, "Hello Word!", anInteger, defaultInt, myString, newString)
}
