package main

import (
	"fmt"
	"io"
)

func Score(output *io.Writer) {
	for i := 3; i > 0; i-- {
		fmt.Fprint(*output, "3")
	}
	fmt.Fprint(*output, "Go!")
}

func main() {
}
