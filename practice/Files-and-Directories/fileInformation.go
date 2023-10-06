package main

import (
	"fmt"
	"os"
)

func checkFileExists(filePath string) bool {
	if _, err := os.Stat(filePath); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func main() {
	stats, err := os.Stat("text.txt")
	if err != nil {
		panic(err)
	}

	//check if file exists
	exists := checkFileExists("text.txt")
	fmt.Println("File exists:", exists)

	// get the file's modification time
	fmt.Println("Modification time:", stats.ModTime())
	fmt.Println("File mode:", stats.Mode())
	fmode := stats.Mode()
	if fmode.IsRegular() {
		fmt.Println("This is a regular file")
	}

	// get the file size
	fmt.Println("File Size:", stats.Size())

	//is this a directory?
	fmt.Println("Is Dir:", stats.IsDir())
}
