package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// read file
	content, err := ioutil.ReadFile("text.txt")
	handleErr(err)

	fmt.Println(string(content))

	const BuffSize = 20
	f, _ := os.Open("text.txt")
	defer f.Close()

	b1 := make([]byte, BuffSize)
	for {
		n, err := f.Read(b1)

		if err != nil {
			if err != io.EOF {
				handleErr(err)
			}
			break
		}
		fmt.Println("bytes read: ", n)
		fmt.Println("Content: ", string(b1[:n]))
	}

}
