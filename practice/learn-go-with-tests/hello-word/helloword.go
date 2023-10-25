package main

import "fmt"

const prefixText = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "Oséias"
	}
	return prefixText + name
}

func main() {
	fmt.Println(Hello("Oséias"))
}
