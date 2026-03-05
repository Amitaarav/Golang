package main

import (
	"fmt"
	"net/url"
)

const myUrl = "https://lco.dev::3000/learn?coursename=reactjs&paymentid=ghbj456ghb"
func main() {
    fmt.Println("Welcome to web server")
	fmt.Println(myUrl)
	//parsing the url
	result, _ := url.Parse(myUrl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)
	qparams := result.Query()
	fmt.Printf("The type of query params are: %T\n", qparams) //url.Values

	fmt.Println(qparams["coursename"])
	fmt.Println(qparams["paymentid"])

	for _, val := range qparams {
		fmt.Println("Param is: ", val)
	}
	// pass reference of url.Values
	partOfUrl := &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "/tutcss",
		RawPath: "user=Amit",
	}
	finalUrl := partOfUrl.String()
	fmt.Println(finalUrl)
}