package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//temporary directory
	tempPath := os.TempDir()
	fmt.Println("Default path", tempPath)

	//temporary file
	tempContent := []byte("This is some temp content for the file")
	tmpFile, err := ioutil.TempFile(tempPath, "tempfile_*.txt")

}
