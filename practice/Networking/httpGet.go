package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	const httpbin = "https://httpbin.org/get"

	//perform a get operation
	resp, err := http.Get(httpbin)
	if err != nil {
		return
	}

	// the caller is responsible for closing the response
	defer resp.Body.Close()

	// we can acess parts of the response to get information
	fmt.Println("status", resp.Status)
	fmt.Println("StatusCode", resp.StatusCode)
	fmt.Println("Proto", resp.Proto)
	fmt.Println("ContentLength", resp.ContentLength)

	// string builder
	var sb strings.Builder

	// read the content and write it to the builder
	content, _ := ioutil.ReadAll(resp.Body)
	bytecount, _ := sb.Write(content)

	// format the output
	fmt.Println(bytecount, sb.String())

}
