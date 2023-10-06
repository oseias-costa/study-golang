package main

import (
	"fmt"
	"net/url"
)

func main() {
	// define a URL
	s := "https://www.example.com:8000/user?username=oseiascosta"

	// // parse function
	result, _ := url.Parse(s)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	// //extract the query components into a value sctruct
	vals := result.Query()
	fmt.Println(vals["username"])

	//create a url from components
	newurl := &url.URL{
		Scheme:   "https",
		Host:     "www.example.com",
		Path:     "/args",
		RawQuery: "x=1&y=2",
	}

	s = newurl.String()
	fmt.Println(s)
	newurl.Host = "joemarini.com"
	fmt.Println(s)

	// create a new values struct and encode it as a query string
	newvals := url.Values{}

	newvals.Add("x", "100")
	newvals.Add("z", "somestr")
	newurl.RawQuery = newvals.Encode()
	s = newurl.String()
	fmt.Println(s)

}
