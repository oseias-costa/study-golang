package main

import (
	"fmt"
	"os"
)

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}

func writeFileData() {
	f, err := os.Create("text2.txt")
	handleErr(err)

	defer f.Close()

	// get the name of the file
	fmt.Println("The file name is", f.Name())

	// write sime string
	f.WriteString("this is some text 22 \n")

	data2 := []byte{'a', 'b', '\n'}
	f.Write(data2)

	f.Sync()
}

func appendFileData(fname string, data string) {

	f, err := os.OpenFile(fname, os.O_APPEND|os.O_WRONLY, 0644)
	handleErr(err)
	defer f.Close()

	_, err = f.Write(data)
}

func main() {
	// //dump some data to a file
	// data1 := []byte("This is some text data \n")
	// ioutil.WriteFile("text.txt", data1, 0644)

	writeFileData()
	appendFileData("text2", "qualquerrrr")
}
