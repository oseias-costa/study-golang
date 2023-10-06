package main

import (
	"fmt"
	"os"
)

func main() {
	tempPath := os.TempDir()
	fmt.Println("Default path", tempPath)

}
