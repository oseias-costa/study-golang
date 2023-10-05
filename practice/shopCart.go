package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

func calculate(input1 string, input2 string, operation string) float64 {
	var result float64

	value1 := convertInputToValue(input1)
	value2 := convertInputToValue(input2)

	switch operation {
	case "+":
		result = addValues(value1, value2)
	case "-":
		result = subtractValues(value1, value2)
	case "*":
		result = multiplyValues(value1, value2)
	case "/":
		result = divideValues(value1, value2)
	default:
		panic("Invalid operation")
	}

	return result
}

func convertInputToValue(input string) float64 {
	value, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		message := fmt.Sprintf("%v must be a number", input)
		panic(message)
	}
	return value
}

func addValues(value1, value2 float64) float64 {
	return value1 + value2
}

func subtractValues(value1, value2 float64) float64 {
	return value1 - value2
}

func multiplyValues(value1, value2 float64) float64 {
	return value1 * value2
}

func divideValues(value1, value2 float64) float64 {
	return value1 / value2
}

func convertToMap(items []string) map[string]float64 {
	result := make(map[string]float64)

	elementValue := 100 / float64(len(items))
	for _, fruit := range items {
		result[fruit] = elementValue
	}
	return result
}

func getCartFromJson(jsonString string) []cartItem {
    var cart []cartItem
    
	err := json.Unmarshal([]byte(jsonString), &cart)
	if err != nil {
		panic("Error reading json string")
	}
	return cart

}