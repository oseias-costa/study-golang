package main

import (
	"fmt"
	"strconv"
	"strings"
)

const showExpectedResult = false
const showHints = false

func calculate(input1 string, input2 string) float64 {
	float1, err := strconv.ParseFloat(strings.TrimSpace(input1), 64)
	if err != nil {
		fmt.Println(err)
		panic("value must be a number")
	}

	float2, err := strconv.ParseFloat(strings.TrimSpace(input2), 64)
	if err != nil {
		fmt.Println(err)
		panic("value must be a number")
	}

	return float1 + float2

}
