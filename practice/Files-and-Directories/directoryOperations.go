package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// create directory
	os.Mkdir("newdir", os.ModePerm)

	// create a deep directory
	os.MkdirAll("path/to/new/dir", os.ModePerm)

	// remove an item
	// defer os.Remove("newdir")

	// remove all
	// defer os.RemoveAll("path")

	// get the current directory
	dir, _ := os.Getwd()
	fmt.Println("current dir is", dir)

	// get the directory of the current process
	exedir, _ := os.Executable()
	fmt.Println("exe dir is", exedir)

	// read the contents of a directory
	contents, _ := ioutil.ReadDir(".")
	for _, fi := range contents {
		fmt.Println(fi.Name(), fi.IsDir())
	}

}
