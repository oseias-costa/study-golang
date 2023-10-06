package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func postRequestTest() {
	const httpbin = "https://httpbin.org/post"

	//post operation using post
	reqBody := strings.NewReader(`
		{
			"field1": "This is field 1",
			"filed2": 250
		}
	`)

	resp, err := http.Post(httpbin, "aplication/json", reqBody)
	if err != nil {
		return
	}

	content, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	fmt.Printf("%s\n", content)

	//post operation using PostForm
	data := url.Values{}
	data.Add("field1", "Field added via values")
	data.Add("field2", "300")
	resp, err = http.PostForm(httpbin, data)
	content, _ = ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	fmt.Printf("%s\n", content)

}

func main() {
	postRequestTest()
}
