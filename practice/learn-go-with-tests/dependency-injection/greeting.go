package main

import (
	"fmt"
	"io"
	"os"
)

func Greeting(writer io.Writer, name string){
	fmt.Printf("Hello, %s", name)
}

func main(){
	Greeting(os.Stdout, "Elodie")
}