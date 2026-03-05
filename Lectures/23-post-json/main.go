package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
    fmt.Println("Welcome to web verb ")
	//PerformGetRequest()
	PerformPostFormRequest();
}

func PerformGetRequest() {
	const myurl = "http://localhost:3000/get"

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content length: ", response.ContentLength)

	var responseString strings.Builder
	content, _ := io.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)

	fmt.Println("Byte count is: ", byteCount)
	fmt.Println(responseString.String())
}

func PerformPostJsonRequest(){
	const myUrl = "http://localhost:3000/post"

	//fake json payload
	requestBody := strings.NewReader(`
	{
	"coursename": "Let's learn Go",
	"price": 299,
	"platform": "LearnCodeOnline.in",
	}
	`)

	response, err := http.Post(myUrl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))

}

func PerformPostFormRequest(){
	const myUrl = "http://localhost:3000/postform"

	data := url.Values{}
	data.Add("firstname", "Rahul")
	data.Add("lastname", "Shetty")
	data.Add("email", "rahulshetty@gmail.com")

	response,error := http.PostForm(myUrl, data);
	if error != nil {
		panic(error)
	}
	defer response.Body.Close()
	content,_ := io.ReadAll(response.Body)
	fmt.Println(string(content))
}
